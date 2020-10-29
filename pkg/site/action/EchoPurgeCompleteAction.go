package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoPurgeCompleteAction struct {
	Action
}

func NewEchoPurgeCompleteAction() *EchoPurgeCompleteAction {
	rca := new(EchoPurgeCompleteAction)
	return rca
}

func (self *EchoPurgeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	messageManager.RemoveMessagesByAreaName(echoTag)

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
