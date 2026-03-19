package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoMsgTwitHandler struct {
	registry *registry.Container
}

func NewEchoMsgTwitHandler(registry *registry.Container) *EchoMsgTwitHandler {
	return &EchoMsgTwitHandler{
		registry: registry,
	}
}

func (self EchoMsgTwitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.registry)
	mapperManager := mapper.RestoreMapperManager(self.registry)
	twitMapper := mapperManager.GetTwitMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Get "echoname" in user request */
	var areaIndex string = r.PathValue("echoname")
	log.Printf("areaIndex = %+v", areaIndex)

	/* Get Echo area by area index */
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName in echoAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Restore message by "echoname" and "msgid" key */
	var msgHash string = r.PathValue("msgid")
	var areaName string = area.GetName()
	origMsg, err3 := echoMapper.GetMessageByHash(areaName, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash in echoMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	newFrom := origMsg.From

	err4 := twitMapper.RegisterTwitByName(newFrom)
	if err4 != nil {
		response := fmt.Sprintf("Fail on RegisterTwitByName in twitMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect */
	echoAddr := urlManager.CreateUrl("/echo/{area_index}").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	http.Redirect(w, r, echoAddr, 303)

}
