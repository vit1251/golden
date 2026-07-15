package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoAreaPurgeHandler struct {
    registry *registry.Container
}

func NewEchoAreaPurgeHandler(registry *registry.Container) *EchoAreaPurgeHandler {
    return &EchoAreaPurgeHandler{
	registry: registry,
    }
}

func (h *EchoAreaPurgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()
    echoMapper := mapperManager.GetEchoMapper()

    areaIndex := r.PathValue("echoname")
    area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    archiveCount, err2 := echoMapper.GetArchivedMessageCount(area.GetName())
    if err2 != nil {
        http.Error(w, err2.Error(), http.StatusInternalServerError)
        return
    }

    data := views.EchoAreaPurgeData{
        Actions: []views.ToolbarAction{
            {Label: "Back", URL: "/echo/" + areaIndex, Icon: "arrow-left"},
        },
        AreaName:        area.GetName(),
        ArchiveMsgCount: archiveCount,
        ActionURL:       "/echo/" + areaIndex + "/purge",
    }

    err := views.Page("Purge area", views.EchoAreaPurgeView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
