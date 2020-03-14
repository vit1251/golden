package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	msg2 "github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/views"
	"log"
	"net/http"
	"path/filepath"
)

type RemoveAction struct {
	Action
}

func NewRemoveAction() *RemoveAction {
	ra:=new(RemoveAction)
	return ra
}

func (self *RemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var messageManager *msg2.MessageManager
	self.Container.Invoke(func(am *area2.AreaManager, mm *msg2.MessageManager) {
		areaManager = am
		messageManager = mm
	})

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	msgHash := vars["msgid"]
	msg, err2 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash on MessageManager: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_msg_remove.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Area", area)
	doc.SetParam("Msg", msg)
	err3 := doc.Render(w)
	if err3 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
