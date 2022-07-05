package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
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

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	/* Parse URL parameters */
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	var areaName string = area.GetName()
	if err1 := echoMapper.RemoveMessagesByAreaName(areaName); err1 != nil {
		log.Printf("err1 = %+v", err1)
	}

	echoAreaMapper.RemoveAreaByName(areaName)

	//
	newLocation := fmt.Sprintf("/echo")
	http.Redirect(w, r, newLocation, 303)

}
