package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoAreaRemoveCompleteAction struct {
	Action
}

func NewEchoRemoveCompleteAction() *EchoAreaRemoveCompleteAction {
	rca := new(EchoAreaRemoveCompleteAction)
	return rca
}

func (self *EchoAreaRemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	if err1 := echoMapper.RemoveMessagesByAreaName(echoTag); err1 != nil {
		log.Printf("err1 = %+v", err1)
	}

	echoAreaMapper.RemoveAreaByName(echoTag)

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
