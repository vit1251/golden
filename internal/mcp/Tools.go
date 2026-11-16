package mcp

import (
    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/registry"
)

func registerTools(srv *mcpsdk.Server, r *registry.Container) {
    registerMailerStatsTool(srv, r)
    registerMailerCheckTool(srv, r)
    registerEchoListTool(srv, r)
    registerEchoReadMessageTool(srv, r)
    registerEchoListMessagesTool(srv, r)
    registerEchoMsgArchiveTool(srv, r)
}
