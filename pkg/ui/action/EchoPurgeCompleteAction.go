package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
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

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	self.Container.Invoke(func(mm *msg.MessageManager) {
		mm.RemoveMessagesByAreaName(echoTag)
	})

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
