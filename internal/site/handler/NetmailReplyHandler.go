package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/registry"
)

type NetmailReplyHandler struct {
	registry *registry.Container
}

func NewNetmailReplyHandler(registry *registry.Container) *NetmailReplyHandler {
	return &NetmailReplyHandler{
		registry: registry,
	}
}

func (self *NetmailReplyHandler) preprocessSubject(origMsg *mapper.NetmailMsg) string {
	compactor := msg.NewSubjectCompactor()
	newSubject := compactor.Compact(origMsg.Subject)
	return newSubject
}

func (self *NetmailReplyHandler) preprocessBody(origMsg *mapper.NetmailMsg) string {

	cmap := msg.NewMessageAuthorParser()
	ma, _ := cmap.Parse(origMsg.From)

	/* Make reply content */
	mtp := msg.NewMessageTextProcessor()
	doc, _ := mtp.Prepare(origMsg.Content)
	newContent := doc.Content()
	log.Printf("reply: orig = %+v", newContent)

	/* Message replay transform */
	mrt := msg.NewMessageReplyTransformer()
	mrt.SetAuthor(ma.QuoteName)
	newContent2 := mrt.Transform(newContent)

	return newContent2
}

func (self *NetmailReplyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	netmailMapper := mapperManager.GetNetmailMapper()
	draftMapper := mapperManager.GetDraftMapper()

	/* Get message hash */
	var msgHash string = r.PathValue("msgid")

	/* Get message */
	origMsg, err3 := netmailMapper.GetMessageByHash(msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if origMsg == nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create new draft message */
	newDraft := mapper.NewDraft()
	newDraft.SetTo(origMsg.GetFrom())
	newDraft.SetToAddr(origMsg.OrigAddr)
	newDraft.SetReply(origMsg.MsgID)

	/* Prepare subject */
	newSubject := self.preprocessSubject(origMsg)
	newDraft.SetSubject(newSubject)

	/* Prepare body */
	newBody := self.preprocessBody(origMsg)
	newDraft.SetBody(newBody)

	/* Store draft message */
	err4 := draftMapper.RegisterNewDraft(*newDraft)
	if err4 != nil {
		response := fmt.Sprintf("Fail in RegisterNewDraft on draftMapper: err = %+v", err4)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/draft/%s/edit", newDraft.GetUUID())
	http.Redirect(w, r, newLocation, 303)

}
