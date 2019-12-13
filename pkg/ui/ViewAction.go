package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	msgProc "github.com/vit1251/golden/pkg/msg"
	"path/filepath"
	"html/template"
	"log"
)

func (self *ViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "view.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	//
	echoTag := params.ByName("name")
	log.Printf("echoTag = %v", echoTag)
	//
	area, err1 := self.Site.app.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)
	//
	msgHeaders, err112 := self.Site.app.MessageBaseReader.GetMessageHeaders(echoTag)
	if (err112 != nil) {
		panic(err112)
	}

	//
	messageId := params.ByName("msgid")
	msg, err2 := self.Site.app.MessageBaseReader.GetMessage(echoTag, messageId)
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
	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.AreaList.Areas
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	outParams["Msg"] = msg
	outParams["Content"] = outDoc
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
