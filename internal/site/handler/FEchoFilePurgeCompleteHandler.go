package handler

import (
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoFilePurgeCompleteHandler struct {
    registry *registry.Container
}

func NewFEchoFilePurgeCompleteHandler(registry *registry.Container) *FEchoFilePurgeCompleteHandler {
    return &FEchoFilePurgeCompleteHandler{
        registry: registry,
    }
}

func (h *FEchoFilePurgeCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileMapper := mapperManager.GetFileMapper()

    areaName := r.PathValue("echoname")
    if err1 := fileMapper.PurgeArchivedFiles(areaName); err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/file/"+areaName, 303)

}
