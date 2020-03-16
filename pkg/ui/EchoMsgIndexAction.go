package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/views"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
	"path/filepath"
)

type EchoMsgIndexAction struct {
	Action
}

func NewEchoMsgIndexAction() *EchoMsgIndexAction {
	ea := new(EchoMsgIndexAction)
	return ea
}

func (self *EchoMsgIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area.AreaManager
	var messageManager *msg.MessageManager
	self.Container.Invoke(func(am *area.AreaManager, mm *msg.MessageManager) {
		areaManager = am
		messageManager = mm
	})

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	newArea, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName where echoTag is %s: err = %+v", echoTag, err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", newArea)

	/* Get message headers */
	msgHeaders, err2 := messageManager.GetMessageHeaders(echoTag)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders where echoTag is %s: err = %+v", echoTag, err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeaders = %+v", msgHeaders)
	for _, msg := range msgHeaders {
		log.Printf("msg = %+v", msg)
	}

	/* Context actions */
	var actions []*widgets.UserAction
	action1 := widgets.NewUserAction()
	action1.Link = fmt.Sprintf("/echo/%s/message/compose", newArea.Name)
	action1.Icon = "/static/img/icon/quote-50.png"
	action1.Label = "Compose"
	actions = append(actions, action1)

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_msg_index.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Actions", actions)
	doc.SetParam("Area", newArea)
	doc.SetParam("Headers", msgHeaders)
	err3 := doc.Render(w)
	if err3 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
