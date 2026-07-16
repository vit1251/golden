package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoAreaCreateHandler struct {
    registry *registry.Container
}

func NewFEchoAreaCreateHandler(registry *registry.Container) *FEchoAreaCreateHandler {
    return &FEchoAreaCreateHandler{
    	registry: registry,
    }
}

func (h *FEchoAreaCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    data := views.FEchoAreaCreateData{
        Actions: []views.ToolbarAction{
            {Label: "Back", URL: "/file", Icon: "arrow-left"},
        },
        ActionURL: "/file/create",
    }
    err := views.Page("Create file area", views.FEchoAreaCreateView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
