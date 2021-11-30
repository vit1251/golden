package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoAreaMarkCompleteAction struct {
	Action
}

func NewEchoAreaMarkCompleteAction() *EchoAreaMarkCompleteAction {
	rca := new(EchoAreaMarkCompleteAction)
	return rca
}

func (self *EchoAreaMarkCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	err1 := echoMapper.MarkAllReadByAreaName(echoTag)
	if err1 != nil {
		log.Printf("Error mark al as read")
	}

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
