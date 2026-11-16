package mcp

import (
    "strings"
    "context"
    "time"

    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type MailerStatsInput struct {
    ErrorsOnly   bool `json:"errors_only,omitempty" jsonschema:"If true, return only sessions whose status indicates an error."`
    WithDataOnly bool `json:"with_data_only,omitempty" jsonschema:"If true, return only sessions with any transferred files (rx>0 or tx>0)."`
}

type MailerSession struct {
    StartedAt   string `json:"started_at"`
    StoppedAt   string `json:"stopped_at"`
    DurationSec int64  `json:"duration_sec"`
    Status      string `json:"status"`
    FilesRx     int    `json:"files_rx"`
    FilesTx     int    `json:"files_tx"`
    IsError     bool   `json:"is_error"`
}

type MailerStatsOutput struct {
    Sessions     []MailerSession `json:"sessions"`
    TotalFilesRx int             `json:"total_files_rx"`
    TotalFilesTx int             `json:"total_files_tx"`
    LastSession  *MailerSession  `json:"last_session,omitempty"`
    Total        int             `json:"total"`
}

func isMailerError(status string) bool {
    s := strings.ToLower(status)
    for _, m := range []string{"error", "fail", "timeout", "abort", "refused", "denied"} {
	if strings.Contains(s, m) {
	    return true
	}
    }
    return false
}

func registerMailerStatsTool(srv *mcpsdk.Server, r *registry.Container) {
    mcpsdk.AddTool(srv, &mcpsdk.Tool{
	Name:        "mailer_stats",
	Description: "List recent mailer sessions with timing, status, and file RX/TX counts. Supports error/data filters.",
    }, func(ctx context.Context, _ *mcpsdk.CallToolRequest, in MailerStatsInput) (*mcpsdk.CallToolResult, MailerStatsOutput, error) {
	_ = ctx
	mm := mapper.RestoreMapperManager(r)
	stats, err := mm.GetStatMailerMapper().GetMailerSummary()
	if err != nil {
	    return nil, MailerStatsOutput{}, err
	}
	out := MailerStatsOutput{Sessions: make([]MailerSession, 0, len(stats))}
        for _, s := range stats {
	    isErr := isMailerError(s.Status)
	    if in.ErrorsOnly && !isErr {
		continue
	    }
	    if in.WithDataOnly && s.FileRXcount == 0 && s.FileTXcount == 0 {
		continue
	    }
	    item := MailerSession{
		StartedAt:   time.UnixMilli(s.SessionStart).UTC().Format(time.RFC3339),
		StoppedAt:   time.UnixMilli(s.SessionStop).UTC().Format(time.RFC3339),
		DurationSec: int64(s.GetDuration() / 1000),
		Status:      s.Status,
		IsError:     isErr,
		FilesRx:     s.FileRXcount,
		FilesTx:     s.FileTXcount,
	    }
	    out.TotalFilesRx += s.FileRXcount
	    out.TotalFilesTx += s.FileTXcount
	    out.Sessions = append(out.Sessions, item)
	}
	if len(out.Sessions) > 0 {
            out.LastSession = &out.Sessions[0]
        }
	out.Total = len(out.Sessions)
	return nil, out, nil
    })
}
