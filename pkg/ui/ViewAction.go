package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	msgProc "github.com/vit1251/golden/pkg/msg"
	"path/filepath"
	"html/template"
	"fmt"
	"log"
)

type ViewAction struct {
	Action
	tmpl     *template.Template
}

func NewViewAction() (*ViewAction) {
	va:=new(ViewAction)

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_msg_view.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	va.tmpl = tmpl

	return va
}

func (self *ViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	webSite := self.Site

	//
	areaManager := webSite.GetAreaManager()
	area, err1 := areaManager.GetAreaByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)

	//
	messageManager := webSite.GetMessageManager()
	msgHeaders, err112 := messageManager.GetMessageHeaders(echoTag)
	if (err112 != nil) {
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
	mr := msgProc.NewMessageTextReader()
	outDoc := mr.Prepare(content)

	/* Update view counter */
	err4 := messageManager.ViewMessageByHash(echoTag, msgHash)
	if err4 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	outParams["Msg"] = msg
	outParams["Content"] = outDoc
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
