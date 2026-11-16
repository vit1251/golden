package mcp

import (
    "context"

    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoListInput struct{}

type EchoAreaInfo struct {
    Name         string `json:"name"`
    AreaIndex    string `json:"area_index"`
    Summary      string `json:"summary"`
    Charset      string `json:"charset"`
    MessageCount int    `json:"message_count"`
    NewCount     int    `json:"new_count"`
}

type EchoListOutput struct {
    Areas []EchoAreaInfo `json:"areas"`
    Total int            `json:"total"`
}

func registerEchoListTool(srv *mcpsdk.Server, r *registry.Container) {
    mcpsdk.AddTool(srv, &mcpsdk.Tool{
	Name:        "echo_list",
	Description: "List echo areas with descriptions, charsets, and message counters.",
    }, func(ctx context.Context, _ *mcpsdk.CallToolRequest, _ EchoListInput) (*mcpsdk.CallToolResult, EchoListOutput, error) {
	_ = ctx
	mm := mapper.RestoreMapperManager(r)
	areas, err := mm.GetEchoAreaMapper().GetAreas()
	if err != nil {
	    return nil, EchoListOutput{}, err
	}

	out := EchoListOutput{Areas: make([]EchoAreaInfo, 0, len(areas))}
	for _, a := range areas {
	    out.Areas = append(out.Areas, EchoAreaInfo{
		Name:         a.GetName(),
		AreaIndex:    a.GetAreaIndex(),
		Summary:      a.GetSummary(),
		Charset:      a.GetCharset(),
		MessageCount: a.GetMessageCount(),
		NewCount:     a.GetNewMessageCount(),
	    })
	}
	out.Total = len(out.Areas)
	return nil, out, nil
    })
}
