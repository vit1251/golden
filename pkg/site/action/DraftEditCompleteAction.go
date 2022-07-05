package action

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/tosser"
	"net/http"
)

type DraftEditCompleteAction struct {
	Action
}

func NewDraftEditCompleteAction() *DraftEditCompleteAction {
	return new(DraftEditCompleteAction)
}

func (self DraftEditCompleteAction) processSaveAction(w http.ResponseWriter, r *http.Request) error {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Get draft */
	newDraft, err1 := self.restoreDraft(r)
	if err1 != nil {
		return err1
	}

	/* Update draft message */
	newTo := r.Form.Get("to")
	newToAddr := r.Form.Get("to_addr")
	newSubject := r.Form.Get("subject")
	newBody := r.Form.Get("body")

	newDraft.SetTo(newTo)
	newDraft.SetToAddr(newToAddr)
	newDraft.SetSubject(newSubject)
	newDraft.SetBody(newBody)

	err2 := draftMapper.UpdateDraft(*newDraft)
	if err2 != nil {
		return err2
	}

	return nil

}

func (self DraftEditCompleteAction) processSendAction(w http.ResponseWriter, r *http.Request) error {

	/* Get draft */
	newDraft, err1 := self.restoreDraft(r)
	if err1 != nil {
		return err1
	}

	if newDraft.IsEchoMail() {
		return self.processConferenceMessage(*newDraft)
	} else {
		return self.processDirectMessage(*newDraft)
	}

}

func (self DraftEditCompleteAction) restoreDraft(r *http.Request) (*mapper.Draft, error) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Get draft message */
	vars := mux.Vars(r)
	draftId := vars["draftid"]

	/* Restore draft index */
	newDraft, err1 := draftMapper.GetDraftByUUID(draftId)
	if err1 != nil {
		return nil, err1
	}
	if newDraft == nil {
		return nil, errors.New("no message exists")
	}

	return newDraft, nil

}

func (self DraftEditCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse POST parameters */
	err1 := r.ParseForm()
	if err1 != nil {
		panic(err1)
	}

	/* Process action */
	actionName := r.Form.Get("action")

	if actionName == "save" {

		/* Process save */
		self.processSaveAction(w, r)

		/* Redirect to drafts */
		newLocation := fmt.Sprintf("/draft")
		http.Redirect(w, r, newLocation, 303)

	} else if actionName == "delivery" {

		/* Process save draft */
		self.processSaveAction(w, r)

		/* Process delivery draft */
		self.processSendAction(w, r)

		/* Process remove draft */
		self.processRemoveAction(w, r)

		/* Redirect to drafts */
		newLocation := fmt.Sprintf("/draft")
		http.Redirect(w, r, newLocation, 303)

	} else if actionName == "remove" {

		self.processRemoveAction(w, r)

		/* Redirect to drafts */
		newLocation := fmt.Sprintf("/draft")
		http.Redirect(w, r, newLocation, 303)

	} else {

		/* Hmmm... */

		/* Redirect to drafts */
		newLocation := fmt.Sprintf("/draft")
		http.Redirect(w, r, newLocation, 303)

	}

}

func (self DraftEditCompleteAction) processConferenceMessage(draft mapper.Draft) error {

	tosserManager := tosser.RestoreTosserManager(self.GetRegistry())

	/* Create message */
	em := tosser.NewEchoMessage()
	em.SetSubject(draft.GetSubject())
	em.SetBody(draft.GetBody())
	em.SetArea(draft.GetArea())
	em.SetTo(draft.GetTo())
	if draft.IsReply() {
		em.SetReply(draft.GetReply())
	}

	/* Delivery conference message */
	err1 := tosserManager.WriteEchoMessage(em)
	if err1 != nil {
		return err1
	}

	/* Register packet */
	//	if err := statMapper.RegisterOutPacket(); err != nil {
	//		log.Printf("Fail on RegisterInPacket: err = %+v", err)
	//	}

	//	if err := statMapper.RegisterOutMessage(); err != nil {
	//		log.Printf("Fail on RegisterOutMessage: err = %+v", err)
	//	}

	return nil

}

func (self DraftEditCompleteAction) processDirectMessage(draft mapper.Draft) error {

	tosserManager := tosser.RestoreTosserManager(self.GetRegistry())

	nm := tosser.NewNetmailMessage()
	nm.SetSubject(draft.GetSubject())
	nm.SetBody(draft.GetBody())
	nm.SetTo(draft.GetTo())
	nm.SetToAddr(draft.GetToAddr())
	if draft.IsReply() {
		nm.SetReply(draft.GetReply())
	}

	/* Delivery direct message */
	err1 := tosserManager.WriteNetmailMessage(nm)
	if err1 != nil {
		return err1
	}

	/* Register packet */
	//	if err := statMapper.RegisterOutPacket(); err != nil {
	//		log.Printf("Fail on RegisterInPacket: err = %+v", err)
	//	}

	//	if err := statMapper.RegisterOutMessage(); err != nil {
	//		log.Printf("Fail on RegisterOutMessage: err = %+v", err)
	//	}

	return nil

}

func (self DraftEditCompleteAction) processRemoveAction(w http.ResponseWriter, r *http.Request) error {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Get draft message */
	vars := mux.Vars(r)
	draftId := vars["draftid"]

	/* Restore draft index */
	err1 := draftMapper.RemoveByUUID(draftId)
	if err1 != nil {
		return err1
	}

	return nil

}
