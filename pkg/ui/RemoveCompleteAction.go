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
	msgid := vars["msgid"]
	//
	self.Site.app.MessageBaseReader.RemoveMessageByHash(echoTag, msgid)
	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}