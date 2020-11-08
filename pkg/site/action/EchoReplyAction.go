package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"log"
	"net/http"
)

type EchoReplyAction struct {
	Action
}

func NewEchoReplyAction() *EchoReplyAction {
	ra := new(EchoReplyAction)
	return ra
}

func (self *EchoReplyAction) preprocessSubject(origMsg msg.Message) string {
	compactor := msg.NewSubjectCompactor()
	newSubject := compactor.Compact(origMsg.Subject)
	return newSubject
}

func (self EchoReplyAction) preprocessBody(origMsg msg.Message) string {

	cmap := msg.NewMessageAuthorParser()
	ma, _ := cmap.Parse(origMsg.From)

	/* Make reply content */
	mtp := msg.NewMessageTextProcessor()
	mtp.Prepare(origMsg.Content)
	newContent := mtp.Content()
	log.Printf("reply: orig = %+v", newContent)

	/* Message replay transform */
	mrt := msg.NewMessageReplyTransformer()
	mrt.SetAuthor(ma.QuoteName)
	newContent2 := mrt.Transform(newContent)

	return newContent2
}

func (self *EchoReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()
	draftMapper := mapperManager.GetDraftMapper()

	/* Get URL params */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]

	/* Get area */
	area, err1 := echoAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail in GetAreaByName on echoAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Get orig message */
	msgHash := vars["msgid"]
	origMsg, err3 := echoMapper.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail in GetMessageByHash on echoMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if origMsg == nil {
		response := fmt.Sprintf("Fail in GetMessageByHash on echoMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create new draft message */
	newDraft := mapper.NewDraft()

	newSubject := self.preprocessSubject(*origMsg)
	newDraft.SetSubject(newSubject)

	newBody := self.preprocessBody(*origMsg)
	newDraft.SetBody(newBody)

	newDraft.SetTo(origMsg.GetFrom())
	newDraft.SetArea(area.GetName())
	newDraft.SetReply(origMsg.GetMsgID())

	/* Store draft message */
	err4 := draftMapper.RegisterNewDraft(*newDraft)
	if err4 != nil {
		response := fmt.Sprintf("Fail in RegisterNewDraft on draftMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect on draft message */
	newLocation := fmt.Sprintf("/draft/%s/edit", newDraft.GetUUID())
	http.Redirect(w, r, newLocation, 303)

}
