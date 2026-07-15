package handler

import (
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoAreaPurgeCompleteHandler struct {
    registry *registry.Container
}

func NewEchoAreaPurgeCompleteHandler(registry *registry.Container) *EchoAreaPurgeCompleteHandler {
    return &EchoAreaPurgeCompleteHandler{
	registry: registry,
    }
}

func (self *EchoAreaPurgeCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(self.registry)
    echoMapper := mapperManager.GetEchoMapper()
    echoAreaMapper := mapperManager.GetEchoAreaMapper()

    var areaIndex string = r.PathValue("echoname")
    area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    var areaName string = area.GetName()
    err2 := echoMapper.PurgeArchivedMessages(areaName)
    if err2 != nil {
        http.Error(w, err2.Error(), http.StatusInternalServerError)
        return
    }

    /* Redirect */
    http.Redirect(w, r, "/echo", 303)

}
