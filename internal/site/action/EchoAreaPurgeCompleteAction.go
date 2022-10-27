package action

import (
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
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

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoMapper := mapperManager.GetEchoMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

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
	err2 := echoMapper.RemoveMessagesByAreaName(areaName)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	echoAddr := urlManager.CreateUrl("/echo").
		Build()
	http.Redirect(w, r, echoAddr, 303)

}
