package handler

import (
    "net/http"
    "github.com/vit1251/golden/internal/site/views"
//    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoAreaCreateHandler struct {
    registry *registry.Container
}

func NewEchoAreaCreateHandler(registry *registry.Container) *EchoAreaCreateHandler {
    return &EchoAreaCreateHandler{
	registry: registry,
    }
}

func (h *EchoAreaCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    data := views.EchoAreaCreateData{
        Actions: []views.ToolbarAction{
            views.ToolbarAction{Label: "Back", URL: "/echo", Icon: "arrow-left"},
        },
        ActionURL: "/echo/create",
    }
    err := views.Page("Create Area", views.EchoAreaCreateView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
