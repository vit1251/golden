package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoAreaUpdateHandler struct {
    registry *registry.Container
}

func NewFEchoAreaUpdateHandler(registry *registry.Container) *FEchoAreaUpdateHandler {
    return &FEchoAreaUpdateHandler{
	registry: registry,
    }
}

func (h *FEchoAreaUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()

    echoName := r.PathValue("echoname")
    area, err1 := fileAreaMapper.GetAreaByName(echoName)
    if err1 != nil || area == nil {
        http.Error(w, "Area not found", http.StatusInternalServerError)
        return
    }

    data := views.FEchoAreaUpdateData{
        Actions: []views.ToolbarAction{
            {Label: "Back",  URL: "/file/" + echoName,            Icon: "arrow-left"},
            {Label: "Purge", URL: "/file/" + echoName + "/purge", Icon: "trash-2"},
        },
        AreaName: area.GetName(),
        Summary:  area.GetSummary(),
        Charset:  area.GetCharset(),
        ActionURL: "/file/" + echoName + "/update",
    }

    err := views.Page("Edit area", views.FEchoAreaUpdateView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
