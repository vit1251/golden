package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/vit1251/golden/pkg/msgapi/squish"
	msgProc "github.com/vit1251/golden/pkg/msg"
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
	var msgBase1 = new(squish.SquishMessageBase)
	msgHeaders, err112 := msgBase1.ReadBase(area.Path)
	if (err112 != nil) {
		panic(err112)
	}

	//
	messageId := params.ByName("msgid")
	var msgId uint64
	msgId, err12 := strconv.ParseUint(messageId, 16, 32)
	log.Printf("err = %v msgid = %d or %x", err12, msgId, msgId)
	//
	var msgBase = new(squish.SquishMessageBase)
	msg, err2 := msgBase.ReadMessage(area.Path, uint32(msgId))
	if (err2 != nil) {
		panic(err2)
	}
	//
	var content string = msg.GetContent()
	//
	mr := msgProc.NewMessageTextReader()
	outDoc := mr.Prepare(content)
	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.config.AreaList.Areas
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	outParams["Msg"] = msg
	outParams["Content"] = outDoc
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
