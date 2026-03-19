package handler

import (
	"fmt"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type NetmailComposeHandler struct {
	registry *registry.Container
}

func NewNetmailComposeHandler(registry *registry.Container) *NetmailComposeHandler {
	return &NetmailComposeHandler{
		registry: registry,
	}
}

func (self *NetmailComposeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Create direct message draft */
	newDraft := mapper.NewDraft()
	//newDraft.SetSubject(subj)
	//newDraft.SetBody(body)
	//newDraft.SetTo(to)
	//newDraft.SetToAddr(to_addr)
	err2 := draftMapper.RegisterNewDraft(*newDraft)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/draft/%s/edit", newDraft.GetUUID())
	http.Redirect(w, r, newLocation, 303)

}
