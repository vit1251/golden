package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoMsgTwitAction struct {
	Action
}

func NewEchoMsgTwitAction() *EchoMsgTwitAction {
	newEchoMsgTwitAction := new(EchoMsgTwitAction)
	return newEchoMsgTwitAction
}

func (self EchoMsgTwitAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := self.restoreUrlManager()
	mapperManager := self.restoreMapperManager()
	twitMapper := mapperManager.GetTwitMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Get "echoname" in user request */
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %+v", areaIndex)

	/* Get Echo area by area index */
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName in echoAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Restore message by "echoname" and "msgid" key */
	msgHash := vars["msgid"]
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
