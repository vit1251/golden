package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/vit1251/golden/pkg/msgapi"
	"path/filepath"
	"html/template"
	"strconv"
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
	area, err1 := self.Site.app.config.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)
	//
	messageId := params.ByName("msgid")
	var msgId uint64
	msgId, err12 := strconv.ParseUint(messageId, 16, 32)
	log.Printf("err = %v msgid = %d or %x", err12, msgId, msgId)
	//
	var msgBase = msgapi.SquishMessageBase{}
	msg, err2 := msgBase.ReadMessage(area.Path, uint32(msgId))
	if (err2 != nil) {
		panic(err2)
	}
	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.config.AreaList.Areas
	outParams["Area"] = area
	outParams["Msg"] = msg
	outParams["Content"] = msg.GetContent()
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
