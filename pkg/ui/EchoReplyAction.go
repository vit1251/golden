package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/views"
	"log"
	"net/http"
	"path/filepath"
)

type EchoReplyAction struct {
	Action
}

func NewEchoReplyAction() *EchoReplyAction {
	ra := new(EchoReplyAction)
	return ra
}

func (self *EchoReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var messageManager *msg.MessageManager
	self.Container.Invoke(func(am *area2.AreaManager, mm *msg.MessageManager) {
		areaManager = am
		messageManager = mm
	})

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Get message area */
	areas, err2 := areaManager.GetAreas()
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	//
	msgHash := vars["msgid"]
	origMsg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Make reply content */
	mtp := msg.NewMessageTextProcessor()
	mtp.Prepare(origMsg.Content)
	mtp.MakeReply()
	newContent := mtp.Content()

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_msg_reply.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Areas", areas)
	doc.SetParam("Area", area)
	doc.SetParam("Msg", origMsg)
	doc.SetParam("Content", newContent)
	err4 := doc.Render(w)
	if err4 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err4)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
