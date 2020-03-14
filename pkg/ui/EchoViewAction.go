package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	msgProc "github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/views"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
	"path/filepath"
)

type EchoViewAction struct {
	Action
}

func NewEchoViewAction() *EchoViewAction {
	va := new(EchoViewAction)
	return va
}

func (self *EchoViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var messageManager *msgProc.MessageManager
	self.Container.Invoke(func(am *area2.AreaManager, mm *msgProc.MessageManager) {
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

	//
	msgHeaders, err112 := messageManager.GetMessageHeaders(echoTag)
	if err112 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	//
	msgHash := vars["msgid"]
	msg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	var content string
	if msg != nil {
		content = msg.GetContent()
	} else {
		content = "!! Unable restore message !!"
	}
	//
	mtp := msgProc.NewMessageTextProcessor()
	err4 := mtp.Prepare(content)
	if err4 != nil {
		response := fmt.Sprintf("Fail on Prepare on MessageTextProcessor")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	outDoc := mtp.HTML()

	/* Update view counter */
	err5 := messageManager.ViewMessageByHash(echoTag, msgHash)
	if err5 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash on messageManager: err = %+v", err5)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create actions */
	var actions []*widgets.UserAction
	action1 := widgets.NewUserAction()
	action1.Link = fmt.Sprintf("/echo/%s//message/%s/reply", area.Name, msg.Hash)
	action1.Icon = "/static/img/icon/quote-50.png"
	action1.Label = "Reply"
	actions = append(actions, action1)
	action2 := 	widgets.NewUserAction()
	action2.Link = fmt.Sprintf("/echo/%s/message/%s/remove", area.Name, msg.Hash)
	action2.Icon = "/static/img/icon/remove-50.png"
	action2.Label = "Delete"
	actions = append(actions, action2)

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_msg_view.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Actions", actions)
	doc.SetParam("Area", area)
	doc.SetParam("Headers", msgHeaders)
	doc.SetParam("Msg", msg)
	doc.SetParam("Content", outDoc)
	err6 := doc.Render(w)
	if err6 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err6)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
