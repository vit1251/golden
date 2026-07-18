package handler

import (
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoFileArchiveHandler struct {
    registry *registry.Container
}

func NewFEchoFileArchiveHandler(registry *registry.Container) *FEchoFileArchiveHandler {
    return &FEchoFileArchiveHandler{
	registry: registry,
    }
}

func (h *FEchoFileArchiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileMapper := mapperManager.GetFileMapper()

    echoName := r.PathValue("echoname")
    file := r.PathValue("file")

    if err1 := fileMapper.ArchiveFileByName(file); err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/file/"+echoName, 303)

}
