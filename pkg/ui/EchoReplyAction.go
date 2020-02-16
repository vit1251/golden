package ui

import (
	"github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/msg"
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"fmt"
	"log"
)

type EchoReplyAction struct {
	Action
}

func NewEchoReplyAction() (*EchoReplyAction) {
	ra := new(EchoReplyAction)
	return ra
}

func (self *EchoReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

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
	areaManager := master.AreaManager
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
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
	messageManager := master.MessageManager
	origMsg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Make reply content */
	mtp := msg.NewMessageTextProcessor()
	mtp.Prepare(origMsg.Content)
	mtp.MakeReply()
	newContent := mtp.Content()

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areas
	outParams["Area"] = area
	outParams["Msg"] = origMsg
	outParams["Content"] = newContent
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
