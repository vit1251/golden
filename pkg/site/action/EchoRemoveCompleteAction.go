package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoRemoveCompleteAction struct {
	Action
}

func NewEchoRemoveCompleteAction() *EchoRemoveCompleteAction {
	rca := new(EchoRemoveCompleteAction)
	return rca
}

func (self *EchoRemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	messageManager := self.restoreMessageManager()
	areaManager := self.restoreAreaManager()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	if err1 := messageManager.RemoveMessagesByAreaName(echoTag); err1 != nil {
		log.Printf("err1 = %+v", err1)
	}

	areaManager.RemoveAreaByName(echoTag)

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
