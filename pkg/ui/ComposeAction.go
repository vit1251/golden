package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"path/filepath"
	"html/template"
	"log"
)

type ComposeAction struct {
	Action
}
type ComposeCompleteAction struct {
	Action
}


func (self *ComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "compose.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	//
	echoTag := params.ByName("name")
	log.Printf("echoTag = %v", echoTag)
	//
	area, err1 := self.Site.app.config.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)
	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.config.AreaList.Areas
	outParams["Area"] = area
	tmpl.ExecuteTemplate(w, "layout", outParams)
}

func (self *ComposeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}