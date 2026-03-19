package handler

import (
	"log"
	"net/http"

	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoMsgRemoveCompleteHandler struct {
	registry *registry.Container
}

func NewEchoMsgRemoveCompleteHandler(registry *registry.Container) *EchoMsgRemoveCompleteHandler {
	return &EchoMsgRemoveCompleteHandler{
		registry: registry,
	}
}

func (self *EchoMsgRemoveCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoMapper := mapperManager.GetEchoMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	urlManager := um.RestoreUrlManager(self.registry)

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	var areaIndex string = r.PathValue("echoname")
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	var areaName string = area.GetName()
	var msgid string = r.PathValue("msgid")
	err2 := echoMapper.RemoveMessageByHash(areaName, msgid)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	areaAddr := urlManager.CreateUrl("/echo/{area_index}").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	http.Redirect(w, r, areaAddr, 303)

}
