package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"log"
)

type RemoveAction struct {
	Action
}
type RemoveCompleteAction struct {
	Action
}

func (self *RemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	//
	area, err1 := self.Site.app.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)

	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.AreaList.Areas
	outParams["Area"] = area
	tmpl.ExecuteTemplate(w, "layout", outParams)
}

func (self *RemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)
	//
	area, err1 := self.Site.app.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)
	//
//	to := r.Form.Get("to")
}
