package mcp

import (
    "context"
    "fmt"
    "strings"
    "time"

    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoListMessagesInput struct {
    Area       string `json:"area" jsonschema:"Echo area name or internal GoldenPoint area hash."`
    Limit      int    `json:"limit,omitempty" jsonschema:"Max messages to return (1..100), default 20."`
    Offset     int    `json:"offset,omitempty" jsonschema:"Offset for pagination, default 0."`
    UnreadOnly bool   `json:"unread_only,omitempty" jsonschema:"If true, return only unread messages."`
}

type EchoMessageHeader struct {
    Hash    string `json:"hash"`
    MsgID   string `json:"msgid,omitempty"`
    Reply   string `json:"reply,omitempty"`
    From    string `json:"from"`
    To      string `json:"to"`
    Subject string `json:"subject"`
    Date    string `json:"date"`
    IsNew   bool   `json:"is_new"`
}

type EchoListMessagesOutput struct {
    Area     string              `json:"area"`
    Total    int                 `json:"total"`
    Limit    int                 `json:"limit"`
    Offset   int                 `json:"offset"`
    Messages []EchoMessageHeader `json:"messages"`
}

func registerEchoListMessagesTool(srv *mcpsdk.Server, r *registry.Container) {
    mcpsdk.AddTool(srv, &mcpsdk.Tool{
	Name:        "echo_list_messages",
	Description: "List message headers in an echo area (paginated).",
    }, func(ctx context.Context, _ *mcpsdk.CallToolRequest, in EchoListMessagesInput) (*mcpsdk.CallToolResult, EchoListMessagesOutput, error) {
	_ = ctx
	limit := in.Limit
	if limit <= 0 { limit = 20 }
	if limit > 100 { limit = 100 }
	offset := in.Offset
	if offset < 0 { offset = 0 }

	mm := mapper.RestoreMapperManager(r)
	eam, em := mm.GetEchoAreaMapper(), mm.GetEchoMapper()

	area, err := eam.GetAreaByName(strings.ToUpper(in.Area))
	if err != nil { return nil, EchoListMessagesOutput{}, err }
	if area == nil {
	    area, err = eam.GetAreaByAreaIndex(in.Area)
	    if err != nil { return nil, EchoListMessagesOutput{}, err }
	}
	if area == nil {
	    return nil, EchoListMessagesOutput{}, fmt.Errorf("area not found: %s", in.Area)
	}

	total, err := em.GetMessageCount(area.GetName())
	if err != nil { return nil, EchoListMessagesOutput{}, err }

	// Желательно: em.GetMessageHeadersPageDesc(area.GetName(), limit, offset)
	headers, err := em.GetMessageHeadersPage(area.GetName(), limit, offset)
	if err != nil { return nil, EchoListMessagesOutput{}, err }

	out := EchoListMessagesOutput{
	    Area: area.GetName(), Total: total, Limit: limit, Offset: offset,
	    Messages: make([]EchoMessageHeader, 0, len(headers)),
	}
	for _, h := range headers {
	    isNew := h.ViewCount == 0
	    if in.UnreadOnly && !isNew { continue }
	    out.Messages = append(out.Messages, EchoMessageHeader{
		Hash: h.Hash, MsgID: h.MsgID, Reply: h.Reply,
		From: h.From, To: h.To, Subject: h.Subject,
		Date:  time.Unix(h.UnixTime, 0).UTC().Format(time.RFC3339),
		IsNew: isNew,
	    })
	}
	return nil, out, nil
    })
}
