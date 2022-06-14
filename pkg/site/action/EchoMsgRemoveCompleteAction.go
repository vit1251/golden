package action

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoMsgRemoveCompleteAction struct {
	Action
}

func NewEchoMsgRemoveCompleteAction() *EchoMsgRemoveCompleteAction {
	rca := new(EchoMsgRemoveCompleteAction)
	return rca
}

func (self *EchoMsgRemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoMapper := mapperManager.GetEchoMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	urlManager := self.restoreUrlManager()

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	var areaName string = area.GetName()
	msgid := vars["msgid"]
	err2 := echoMapper.RemoveMessageByHash(areaName, msgid)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	areaAddr := urlManager.CreateUrl("/echo/{area_index}").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	http.Redirect(w, r, areaAddr, 303)

}
