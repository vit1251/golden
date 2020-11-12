package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoAreaPurgeCompleteAction struct {
	Action
}

func NewEchoAreaPurgeCompleteAction() *EchoAreaPurgeCompleteAction {
	rca := new(EchoAreaPurgeCompleteAction)
	return rca
}

func (self *EchoAreaPurgeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
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

	echoMapper.RemoveMessagesByAreaName(echoTag)

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
