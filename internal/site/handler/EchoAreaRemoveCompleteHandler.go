package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoAreaRemoveCompleteHandler struct {
	registry *registry.Container
}

func NewEchoRemoveCompleteHandler(registry *registry.Container) *EchoAreaRemoveCompleteHandler {
	return &EchoAreaRemoveCompleteHandler{
		registry: registry,
	}
}

func (self *EchoAreaRemoveCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

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
	if err1 := echoMapper.RemoveMessagesByAreaName(areaName); err1 != nil {
		log.Printf("err1 = %+v", err1)
	}

	echoAreaMapper.RemoveAreaByName(areaName)

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
