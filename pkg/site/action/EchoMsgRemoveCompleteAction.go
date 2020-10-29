package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoMsgRemoveCompleteAction struct {
	Action
}

func NewEchoMsgRemoveCompleteAction() *EchoMsgRemoveCompleteAction {
	rca := new(EchoMsgRemoveCompleteAction)
	return rca
}

func (self *EchoMsgRemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	messageManager := self.restoreMessageManager()

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
	messageManager.RemoveMessageByHash(echoTag, msgid)

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)

}
