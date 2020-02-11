package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"fmt"
	"log"
)

type ReplyAction struct {
	Action
}

func NewReplyAction() (*ReplyAction) {
	ra := new(ReplyAction)
	return ra
}

func (self *ReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_msg_reply.tmpl")
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

	/* Get message area */
	areas, err2 := areaManager.GetAreas()
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	//
	msgHash := vars["msgid"]
	messageManager := webSite.GetMessageManager()
	msg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if (err3 != nil) {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = areas
	outParams["Area"] = area
	outParams["Msg"] = msg
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
