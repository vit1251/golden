package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"fmt"
	"log"
)

type EchoComposeAction struct {
	Action
	tmpl *template.Template
}

func NewEchoComposeAction() (*EchoComposeAction) {
	ca := new(EchoComposeAction)

	//
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_msg_compose.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	ca.tmpl = tmpl

	return ca
}

func (self *EchoComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Search area */
	webSite := self.Site

	/* Search echo area */
	areaManager := webSite.GetAreaManager()
	area, err1 := areaManager.GetAreaByName(echoTag)
	if (err1 != nil) {
		response := fmt.Sprintf("Fail on GetAreaByName")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %v", area)

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areas
	outParams["Area"] = area
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
