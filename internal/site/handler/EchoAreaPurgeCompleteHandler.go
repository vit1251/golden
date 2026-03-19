package handler

import (
	"log"
	"net/http"

	"github.com/vit1251/golden/internal/um"
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

	urlManager := um.RestoreUrlManager(self.registry)
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoMapper := mapperManager.GetEchoMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	/* Parse URL parameters */
	var areaIndex string = r.PathValue("echoname")
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	var areaName string = area.GetName()
	err2 := echoMapper.RemoveMessagesByAreaName(areaName)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	echoAddr := urlManager.CreateUrl("/echo").
		Build()
	http.Redirect(w, r, echoAddr, 303)

}
