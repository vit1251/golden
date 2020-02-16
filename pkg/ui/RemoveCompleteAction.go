package ui

import (
	"github.com/vit1251/golden/pkg/common"
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

	master := common.GetMaster()

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
	messageManager := master.MessageManager
	messageManager.RemoveMessageByHash(echoTag, msgid)

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
