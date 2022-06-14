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
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("echoTag = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	var areaName string = area.GetName()
	err2 := echoMapper.MarkAllReadByAreaName(areaName)
	if err2 != nil {
		panic(err2)
	}

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
