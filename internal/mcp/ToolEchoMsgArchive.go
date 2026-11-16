package mcp

import (
    "context"
    "fmt"
    "strings"

    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoMsgArchiveInput struct {
    Area string `json:"area" jsonschema:"Echo area name or internal GoldenPoint area hash."`
    Hash string `json:"hash" jsonschema:"Internal GoldenPoint message hash."`
}

type EchoMsgArchiveOutput struct {
    Archived bool   `json:"archived"`
    Area     string `json:"area"`
    Hash     string `json:"hash"`
}

func registerEchoMsgArchiveTool(srv *mcpsdk.Server, r *registry.Container) {
    mcpsdk.AddTool(srv, &mcpsdk.Tool{
	Name:        "echo_msg_archive",
	Description: "Archive an echo message by area and GoldenPoint hash (marks it archived; it will no longer appear in message lists).",
    }, func(ctx context.Context, _ *mcpsdk.CallToolRequest, in EchoMsgArchiveInput) (*mcpsdk.CallToolResult, EchoMsgArchiveOutput, error) {
	_ = ctx
	mm := mapper.RestoreMapperManager(r)
	eam := mm.GetEchoAreaMapper()
	em := mm.GetEchoMapper()

	area, err := eam.GetAreaByName(strings.ToUpper(in.Area))
	if err != nil {
	    return nil, EchoMsgArchiveOutput{}, err
	}
	if area == nil {
	    area, err = eam.GetAreaByAreaIndex(in.Area)
	    if err != nil {
		return nil, EchoMsgArchiveOutput{}, err
	    }
	}
	if area == nil {
	    return nil, EchoMsgArchiveOutput{}, fmt.Errorf("area not found: %s", in.Area)
	}

	msg, err := em.GetMessageByHash(area.GetName(), in.Hash)
	if err != nil {
	    return nil, EchoMsgArchiveOutput{}, err
	}
	if msg == nil {
	    return nil, EchoMsgArchiveOutput{}, fmt.Errorf("message not found: %s", in.Hash)
	}

	if err := em.ArchiveMessageByHash(area.GetName(), in.Hash); err != nil {
	    return nil, EchoMsgArchiveOutput{}, err
	}

	return nil, EchoMsgArchiveOutput{
	    Archived: true,
	    Area:     area.GetName(),
	    Hash:     in.Hash,
	}, nil
    })
}