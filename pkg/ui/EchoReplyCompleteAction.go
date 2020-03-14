package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/tosser"
	"log"
	"net/http"
)

type ReplyCompleteAction struct {
	Action
}

func NewReplyCompleteAction() (*ReplyCompleteAction) {
	rca := new(ReplyCompleteAction)
	return rca
}

func (self *ReplyCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var messageManager *msg.MessageManager
	var tosserManager *tosser.TosserManager
	self.Container.Invoke(func(am *area2.AreaManager, mm *msg.MessageManager, tm *tosser.TosserManager) {
		areaManager = am
		messageManager = mm
		tosserManager = tm
	})

	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	//
	vars := mux.Vars(r)

	/* Recover area */
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Recover message */
	msgHash := vars["msgid"]
	origMsg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	//
	to := r.Form.Get("to")
	subj := r.Form.Get("subject")
	body := r.Form.Get("body")
	log.Printf("to = %s subj = %s body = %s", to, subj, body)

	/* Create message */
	em := tosserManager.NewEchoMessage()
	em.Subject = subj
	em.Body = body
	em.AreaName = area.Name
	em.To = to
	em.Reply = origMsg.MsgID

	/* Delivery message */
	err4 := tosserManager.WriteEchoMessage(em)
	if err4 != nil {
		panic(err4)
	}

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
