package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type EchoComposeAction struct {
	Action
}

func NewEchoComposeAction() *EchoComposeAction {
	ca := new(EchoComposeAction)
	return ca
}

func (self *EchoComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	draftMapper := mapperManager.GetDraftMapper()

	/* Get URL params */
	vars := mux.Vars(r)
	areaName := vars["echoname"]

	/* Get area */
	area, err1 := echoAreaMapper.GetAreaByName(areaName)
	if err1 != nil {
		response := fmt.Sprintf("Fail in GetAreaByName on echoAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create new draft */
	newDraft := mapper.NewDraft()
	newDraft.SetArea(area.GetName())
	err2 := draftMapper.RegisterNewDraft(*newDraft)
	if err2 != nil {
		response := fmt.Sprintf("Fail in RegisterNewDraft on draftMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect to new draft message */
	newLocation := fmt.Sprintf("/draft/%s/edit", newDraft.GetUUID())
	http.Redirect(w, r, newLocation, 303)

}
