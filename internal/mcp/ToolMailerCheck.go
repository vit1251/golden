package mcp

import (
    "context"

    "github.com/vit1251/golden/pkg/eventbus"
    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/registry"
)

type MailerCheckInput struct{}

type MailerCheckOutput struct {
    Triggered bool   `json:"triggered"`
    Message   string `json:"message"`
}

func registerMailerCheckTool(srv *mcpsdk.Server, r *registry.Container) {
    mcpsdk.AddTool(srv, &mcpsdk.Tool{
	Name:        "mailer_check",
	Description: "Request an immediate mailer session to send/receive mail out of schedule.",
    }, func(ctx context.Context, _ *mcpsdk.CallToolRequest, _ MailerCheckInput) (*mcpsdk.CallToolResult, MailerCheckOutput, error) {
	_ = ctx

	go func() {
	    eventBus := eventbus.RestoreEventBus(r)
	    evt := eventBus.CreateEvent("Mailer")
	    eventBus.FireEvent(evt)
	}()

	return nil, MailerCheckOutput{
	    Triggered: true,
	    Message:   "Mailer session requested.",
	}, nil
    })
}
