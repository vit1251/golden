package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/tosser"
)

type DraftEditCompleteHandler struct {
	registry *registry.Container
}

func NewDraftEditCompleteHandler(registry *registry.Container) *DraftEditCompleteHandler {
	return &DraftEditCompleteHandler{
		registry: registry,
	}
}

func (h *DraftEditCompleteHandler) processSaveHandler(w http.ResponseWriter, r *http.Request) error {

	mapperManager := mapper.RestoreMapperManager(h.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Get draft */
	newDraft, err1 := h.restoreDraft(r)
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

func (h *DraftEditCompleteHandler) processSendHandler(w http.ResponseWriter, r *http.Request) error {

	/* Get draft */
	newDraft, err1 := h.restoreDraft(r)
	if err1 != nil {
		return err1
	}

	if newDraft.IsEchoMail() {
		return h.processConferenceMessage(*newDraft)
	} else {
		return h.processDirectMessage(*newDraft)
	}

}

func (h *DraftEditCompleteHandler) restoreDraft(r *http.Request) (*mapper.Draft, error) {

	mapperManager := mapper.RestoreMapperManager(h.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Get draft message */
	var draftId string = r.PathValue("draftid")

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

func (h *DraftEditCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse POST parameters */
	err1 := r.ParseForm()
	if err1 != nil {
		panic(err1)
	}

	/* Process action */
	actionName := r.Form.Get("action")

	if actionName == "save" {

		/* Process save */
		h.processSaveHandler(w, r)

		/* Redirect to drafts */
		newLocation := fmt.Sprintf("/draft")
		http.Redirect(w, r, newLocation, 303)

	} else if actionName == "delivery" {

		/* Process save draft */
		h.processSaveHandler(w, r)

		/* Process delivery draft */
		h.processSendHandler(w, r)

		/* Process remove draft */
		h.processRemoveHandler(w, r)

		/* Redirect to drafts */
		newLocation := fmt.Sprintf("/draft")
		http.Redirect(w, r, newLocation, 303)

	} else if actionName == "remove" {

		h.processRemoveHandler(w, r)

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

func (h *DraftEditCompleteHandler) processConferenceMessage(draft mapper.Draft) error {

	tosserManager := tosser.RestoreTosserManager(h.registry)

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
	//    if err := statMapper.RegisterOutPacket(); err != nil {
	//        log.Printf("Fail on RegisterInPacket: err = %+v", err)
	//    }

	//    if err := statMapper.RegisterOutMessage(); err != nil {
	//        log.Printf("Fail on RegisterOutMessage: err = %+v", err)
	//    }

	return nil

}

func (h *DraftEditCompleteHandler) processDirectMessage(draft mapper.Draft) error {

	tosserManager := tosser.RestoreTosserManager(h.registry)

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
	//    if err := statMapper.RegisterOutPacket(); err != nil {
	//        log.Printf("Fail on RegisterInPacket: err = %+v", err)
	//    }

	//    if err := statMapper.RegisterOutMessage(); err != nil {
	//        log.Printf("Fail on RegisterOutMessage: err = %+v", err)
	//    }

	return nil

}

func (h *DraftEditCompleteHandler) processRemoveHandler(w http.ResponseWriter, r *http.Request) error {

	mapperManager := mapper.RestoreMapperManager(h.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Get draft message */
	var draftId string = r.PathValue("draftid")

	/* Restore draft index */
	err1 := draftMapper.RemoveByUUID(draftId)
	if err1 != nil {
		return err1
	}

	return nil

}
