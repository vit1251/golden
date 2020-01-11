package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"log"
)

type ComposeAction struct {
	Action
}

func (self *ComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "compose.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Search area */
	webSite := self.Site
	areaManager := webSite.GetAreaManager()
	area, err1 := areaManager.GetAreaByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)

	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = areaManager.GetAreas()
	outParams["Area"] = area
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
