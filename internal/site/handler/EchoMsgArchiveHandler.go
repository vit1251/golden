package handler

import (
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoMsgArchiveHandler struct {
    registry *registry.Container
}

func NewEchoMsgArchiveHandler(registry *registry.Container) *EchoMsgArchiveHandler {
    return &EchoMsgArchiveHandler{
	registry: registry,
    }
}

func (h *EchoMsgArchiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()
    echoMapper := mapperManager.GetEchoMapper()

    areaIndex := r.PathValue("echoname")
    area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    msgid := r.PathValue("msgid")
    areaName := area.GetName() // TODO - заменить на areaIndex
    if err2 := echoMapper.ArchiveMessageByHash(areaName, msgid); err2 != nil {
        http.Error(w, err2.Error(), http.StatusInternalServerError)
	return
    }

    http.Redirect(w, r, "/echo/" + areaIndex, 303)

}
