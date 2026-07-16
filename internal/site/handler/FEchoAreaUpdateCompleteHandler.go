package handler

import (
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoAreaUpdateCompleteHandler struct {
    registry *registry.Container
}

func NewFEchoAreaUpdateCompleteHandler(registry *registry.Container) *FEchoAreaUpdateCompleteHandler {
    return &FEchoAreaUpdateCompleteHandler{
        registry: registry,
    }
}

func (h *FEchoAreaUpdateCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()

    areaName := r.PathValue("echoname")
    area, err := fileAreaMapper.GetAreaByName(areaName)
    if err != nil || area == nil {
        http.Error(w, "Area not found", http.StatusInternalServerError)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    area.SetSummary(r.PostForm.Get("summary"))
    area.SetCharset(r.PostForm.Get("charset"))

    if err := fileAreaMapper.UpdateArea(area); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    http.Redirect(w, r, "/file/"+areaName, 303)

}
