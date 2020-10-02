package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
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

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	self.Container.Invoke(func(am *msg.AreaManager, mm *msg.MessageManager) {
		err1 := mm.RemoveMessagesByAreaName(echoTag)
		log.Printf("err1 = %+v", err1)
		am.RemoveAreaByName(echoTag)
	})

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
