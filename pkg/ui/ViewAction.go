package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	msgProc "github.com/vit1251/golden/pkg/msg"
	"path/filepath"
	"html/template"
	"log"
)

type ViewAction struct {
	Action
}

func (self *ViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "view.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

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
		panic(err112)
	}

	//
	msgHash := vars["msgid"]
	msg, err2 := messageManager.GetMessageByHash(echoTag, msgHash)
	if (err2 != nil) {
		panic(err2)
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

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areaManager.GetAreas()
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	outParams["Msg"] = msg
	outParams["Content"] = outDoc
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
