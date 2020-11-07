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

func (self *NetmailRemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	netmailMapper := mapperManager.GetNetmailMapper()

	//
	vars := mux.Vars(r)

	//
	msgHash := vars["msgid"]
	_, err3 := netmailMapper.GetMessageByHash(msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	log.Printf("msg = %s\n", msgHash)

	if err := netmailMapper.RemoveMessageByHash(msgHash) ; err != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/netmail")
	http.Redirect(w, r, newLocation, 303)

}