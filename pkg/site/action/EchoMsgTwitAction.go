package action

import (
	"fmt"
	"github.com/gorilla/mux"
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

	mapperManager := self.restoreMapperManager()
	twitMapper := mapperManager.GetTwitMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Get "echoname" in user request */
	vars := mux.Vars(r)

	/* Restore area by "echoname" key */
	echoTag := vars["echoname"]
	//log.Printf("echoTag = %+v", echoTag)
	area, err1 := echoAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName in echoAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Restore message by "echoname" and "msgid" key */
	msgHash := vars["msgid"]
	//log.Printf("msgid = %+v", msgid)
	origMsg, err3 := echoMapper.GetMessageByHash(echoTag, msgHash)
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
	areaName := area.GetName()
	newLocation := fmt.Sprintf("/echo/%s", areaName)
	http.Redirect(w, r, newLocation, 303)

}