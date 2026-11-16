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

type EchoReadMessageInput struct {
    Area     string `json:"area" jsonschema:"Echo area name or internal GoldenPoint area hash."`
    Hash     string `json:"hash" jsonschema:"Internal GoldenPoint message hash (stable internal message identifier)."`
}

type EchoReadMessageOutput struct {
    Area     string `json:"area"`
    Hash     string `json:"hash"`
    MsgID    string `json:"msgid,omitempty"`
    Reply    string `json:"reply,omitempty"`
    From     string `json:"from"`
    FromAddr string `json:"from_addr,omitempty"`
    To       string `json:"to"`
    Subject  string `json:"subject"`
    Date     string `json:"date"`
    Body     string `json:"body"`
}

func registerEchoReadMessageTool(srv *mcpsdk.Server, r *registry.Container) {
    mcpsdk.AddTool(srv, &mcpsdk.Tool{
	Name:        "echo_read_message",
	Description: "Read a single echo message by area and GoldenPoint hash. Does not mark the message as read.",
    }, func(ctx context.Context, _ *mcpsdk.CallToolRequest, in EchoReadMessageInput) (*mcpsdk.CallToolResult, EchoReadMessageOutput, error) {
	_ = ctx
	mm := mapper.RestoreMapperManager(r)
	eam := mm.GetEchoAreaMapper()
	em := mm.GetEchoMapper()

	area, err := eam.GetAreaByName(strings.ToUpper(in.Area))
	if err != nil {
	    return nil, EchoReadMessageOutput{}, err
	}
	if area == nil {
	    area, err = eam.GetAreaByAreaIndex(in.Area)
	    if err != nil {
		return nil, EchoReadMessageOutput{}, err
	    }
	}
	if area == nil {
	    return nil, EchoReadMessageOutput{}, fmt.Errorf("area not found: %s", in.Area)
	}

	msg, err := em.GetMessageByHash(area.GetName(), in.Hash)
	if err != nil {
	    return nil, EchoReadMessageOutput{}, err
	}
	if msg == nil {
	    return nil, EchoReadMessageOutput{}, fmt.Errorf("message not found: %s", in.Hash)
	}

	return nil, EchoReadMessageOutput{
	    Area:     area.GetName(),
	    Hash:     msg.Hash,
	    MsgID:    msg.MsgID,
	    Reply:    msg.Reply,
	    From:     msg.From,
	    FromAddr: msg.FromAddr,
	    To:       msg.To,
	    Subject:  msg.Subject,
	    Date:     time.Unix(msg.UnixTime, 0).UTC().Format(time.RFC3339),
	    Body:     msg.Content,
	}, nil
    })
}
