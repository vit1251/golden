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

type EchoUpdateAction struct {
	Action
}

func NewEchoUpdateAction() *EchoUpdateAction {
	ea := new(EchoUpdateAction)
	return ea
}

func (self *EchoUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	self.Container.Invoke(func(am *area2.AreaManager) {
		areaManager = am
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

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_update.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Area", area)
	err2 := doc.Render(w)
	if err2 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
