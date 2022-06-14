package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"log"
	"net/http"
)

type EchoMsgComposeAction struct {
	Action
}

func NewEchoMsgComposeAction() *EchoMsgComposeAction {
	ca := new(EchoMsgComposeAction)
	return ca
}

func (self *EchoMsgComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := self.restoreUrlManager()
	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	draftMapper := mapperManager.GetDraftMapper()

	/* Get URL params */
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	/* Get area */
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

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
	draftAddr := urlManager.CreateUrl("/draft/{draft_index}/edit").
		SetParam("draft_index", newDraft.GetUUID()).
		Build()
	http.Redirect(w, r, draftAddr, 303)

}
