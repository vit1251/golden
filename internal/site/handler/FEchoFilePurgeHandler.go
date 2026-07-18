package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoFilePurgeHandler struct {
    registry *registry.Container
}

func NewFEchoFilePurgeHandler(registry *registry.Container) *FEchoFilePurgeHandler {
    return &FEchoFilePurgeHandler{
        registry: registry,
    }
}

func (h *FEchoFilePurgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()
    fileMapper := mapperManager.GetFileMapper()

    areaName := r.PathValue("echoname")
    area, err := fileAreaMapper.GetAreaByName(areaName)
    if err != nil || area == nil {
        http.Error(w, "Area not found", http.StatusInternalServerError)
        return
    }

    archivedCount, err := fileMapper.GetArchivedFileCount(areaName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := views.FEchoFilePurgeData{
        Actions: []views.ToolbarAction{
            {Label: "Back", URL: "/file/" + areaName, Icon: "arrow-left"},
        },
        AreaName:      area.GetName(),
        ArchivedCount: archivedCount,
        ActionURL:     "/file/" + areaName + "/purge",
    }

    err = views.Page("Purge file area", views.FEchoFilePurgeView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
