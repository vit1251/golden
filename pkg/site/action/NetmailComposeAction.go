package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type NetmailComposeAction struct {
	Action
}

func NewNetmailComposeAction() *NetmailComposeAction {
	newNetmailComposeAction := new(NetmailComposeAction)
	return newNetmailComposeAction
}

func (self *NetmailComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
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
