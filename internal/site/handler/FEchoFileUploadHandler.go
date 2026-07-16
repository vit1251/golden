package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/config"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoFileUploadHandler struct {
    registry *registry.Container
}

func NewFEchoFileUploadHandler(registry *registry.Container) *FEchoFileUploadHandler {
    return &FEchoFileUploadHandler{
	registry: registry,
    }
}

func (h *FEchoFileUploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    configManager := config.RestoreConfigManager(h.registry)
    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()

    /* Get BOSS address */
    newConfig := configManager.GetConfig()

    var echoName string = r.PathValue("echoname")
    area, err1 := fileAreaMapper.GetAreaByName(echoName)
    if err1 != nil || area == nil {
        http.Error(w, "Area not found", http.StatusInternalServerError)
    	return
    }

    data := views.FEchoFileUploadData{
        Actions: []views.ToolbarAction{
            {Label: "Back", URL: "/file/" + echoName, Icon: "arrow-left"},
        },
        ActionURL: "/file/" + echoName + "/upload",
        To:        newConfig.Main.Link,
        AreaName:  area.GetName(),
    }

    err := views.Page("Upload file", views.FEchoFileUploadView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
