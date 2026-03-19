package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoMsgComposeHandler struct {
	registry *registry.Container
}

func NewEchoMsgComposeHandler(registry *registry.Container) *EchoMsgComposeHandler {
	return &EchoMsgComposeHandler{
		registry: registry,
	}
}

func (self *EchoMsgComposeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.registry)
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	draftMapper := mapperManager.GetDraftMapper()

	/* Get URL params */
	var areaIndex string = r.PathValue("echoname")
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
