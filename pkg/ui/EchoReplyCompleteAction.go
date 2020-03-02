package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/common"
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

	master := common.GetMaster()

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
	areaManager := master.AreaManager
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Recover message */
	msgHash := vars["msgid"]
	messageManager := master.MessageManager
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
	log.Printf("to = %s subj = %s body = %s reply = %s", to, subj, body)

	/* Create message */
	em := master.TosserManager.NewEchoMessage()
	em.Subject = subj
	em.Body = body
	em.AreaName = area.Name
	em.To = to
	em.Reply = origMsg.MsgID

	/* Delivery message */
	err4 := master.TosserManager.WriteEchoMessage(em)
	if err4 != nil {
		panic(err4)
	}

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
