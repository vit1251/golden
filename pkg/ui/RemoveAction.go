package ui

import (
	"github.com/vit1251/golden/pkg/common"
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"log"
)

type RemoveAction struct {
	Action
}

func NewRemoveAction() (*RemoveAction) {
	ra:=new(RemoveAction)
	return ra
}

func (self *RemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_msg_remove.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	areaManager := master.AreaManager
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %v", area)

	//
	msgHash := vars["msgid"]
	messageManager := master.MessageManager
	msg, err2 := messageManager.GetMessageByHash(echoTag, msgHash)
	if (err2 != nil) {
		panic(err2)
	}

	//
	outParams := make(map[string]interface{})
	outParams["Area"] = area
	outParams["Msg"] = msg
	tmpl.ExecuteTemplate(w, "layout", outParams)

}
