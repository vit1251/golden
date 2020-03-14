package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/ui/views"
	"log"
	"net/http"
	"path/filepath"
)

type EchoComposeAction struct {
	Action
}

func NewEchoComposeAction() *EchoComposeAction {
	ca := new(EchoComposeAction)
	return ca
}

func (self *EchoComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	self.Container.Invoke(func(am *area2.AreaManager) {
		areaManager = am
	})

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Search echo area */

	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_msg_compose.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Areas", areas)
	doc.SetParam("Area", area)
	err2 := doc.Render(w)
	if err2 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
