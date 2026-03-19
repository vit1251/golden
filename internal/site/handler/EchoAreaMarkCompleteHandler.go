package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoAreaMarkCompleteHandler struct {
	registry *registry.Container
}

func NewEchoAreaMarkCompleteHandler(registry *registry.Container) *EchoAreaMarkCompleteHandler {
	return &EchoAreaMarkCompleteHandler{
		registry: registry,
	}
}

func (self *EchoAreaMarkCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	var areaIndex string = r.PathValue("echoname")
	log.Printf("echoTag = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	var areaName string = area.GetName()
	err2 := echoMapper.MarkAllReadByAreaName(areaName)
	if err2 != nil {
		panic(err2)
	}

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
