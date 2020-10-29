package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NetmailRemoveAction struct {
	Action
}

func NewNetmailRemoveAction() *NetmailRemoveAction {
	return new(NetmailRemoveAction)
}

func (self *NetmailRemoveAction)  ServeHTTP(w http.ResponseWriter, r *http.Request) {

	netmailManager := self.restoreNetmailManager()

	//
	vars := mux.Vars(r)

	//
	msgHash := vars["msgid"]
	//origMsg, err3 := netmailManager.GetMessageByHash(msgHash)
	//if err3 != nil {
	//	response := fmt.Sprintf("Fail on GetMessageByHash")
	//	http.Error(w, response, http.StatusInternalServerError)
	//	return
	//}

	log.Printf("msg = %s\n", msgHash)

	if err := netmailManager.RemoveMessageByHash(msgHash) ; err != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/netmail")
	http.Redirect(w, r, newLocation, 303)

}