package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
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

	var messageManager *msg.MessageManager
	self.Container.Invoke(func(mm *msg.MessageManager) {
		messageManager = mm
	})

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
