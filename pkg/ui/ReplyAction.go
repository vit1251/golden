package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"log"
)

type ReplyAction struct {
	Action
}

func (self *ReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "reply.tmpl")
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
	msgHash := vars["msgid"]
	messageManager := webSite.GetMessageManager()
	msg, err2 := messageManager.GetMessageByHash(echoTag, msgHash)
	if (err2 != nil) {
		panic(err2)
	}

	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = areaManager.GetAreas()
	outParams["Area"] = area
	outParams["Msg"] = msg
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
