package mcp

import (
    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/registry"
)

func registerTools(srv *mcpsdk.Server, r *registry.Container) {
    _ = srv
    _ = r

    // TODO: register tools here, for example:
    // mcpsdk.AddTool(srv, &mcpsdk.Tool{
    // 	Name:        "list_echoareas",
    // 	Description: "List echo areas.",
    // }, handlerFunc)
}
