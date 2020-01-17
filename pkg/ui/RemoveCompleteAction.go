package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

type RemoveCompleteAction struct {
	Action
}

func NewRemoveCompleteAction() (*RemoveCompleteAction) {
	rca := new(RemoveCompleteAction)
	return rca
}

func (self *RemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	webSite := self.Site

	//
	msgid := vars["msgid"]
	messageManager := webSite.GetMessageManager()
	messageManager.RemoveMessageByHash(echoTag, msgid)

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
