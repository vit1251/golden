package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/common"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type EchoAction struct {
	Action
	tmpl  *template.Template
}

func NewEchoAction() (*EchoAction) {
	ea := new(EchoAction)

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_msg_index.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	ea.tmpl = tmpl

	return ea
}

func (self *EchoAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get area manager */
	areaManager := master.AreaManager

	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName where echoTag is %s: err = %+v", echoTag, err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	/* Get message headers */
	messageManager := master.MessageManager
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
	var actions []*UserAction
	action1 := NewUserAction()
	action1.Link = fmt.Sprintf("/echo/%s/message/compose", area.Name)
	action1.Icon = "/static/img/icon/quote-50.png"
	action1.Label = "Compose"
	actions = append(actions, action1)

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Actions"] = actions
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	self.tmpl.ExecuteTemplate(w, "layout", outParams)

}
