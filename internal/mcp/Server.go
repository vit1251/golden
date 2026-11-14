package mcp

import (
    "net/http"
    "runtime/debug"

    mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/vit1251/golden/pkg/registry"
)

func NewHandler(r *registry.Container) http.Handler {
    srv := mcpsdk.NewServer(&mcpsdk.Implementation{
	Name:    "golden-point",
	Version: serverVersion(),
    }, nil)

    registerTools(srv, r)

    return mcpsdk.NewStreamableHTTPHandler(func(_ *http.Request) *mcpsdk.Server {
	return srv
    }, &mcpsdk.StreamableHTTPOptions{
	Stateless:    true,
	JSONResponse: true,
	// Logger: slog.Default(),
    })
}

func serverVersion() string {
    if info, ok := debug.ReadBuildInfo(); ok {
	if v := info.Main.Version; v != "" && v != "(devel)" {
	    return v
	}
    }
    return "dev"
}
