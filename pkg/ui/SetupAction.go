package ui

import (
	"net/http"
//	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"log"
)

type SetupAction struct {
	Action
}

func NewSetupAction() (*SetupAction) {
	sa := new(SetupAction)
	return sa
}

func (self *SetupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "setup.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

	//
	webSite := self.Site

	/* Setup manager operation */
	setupManager := webSite.GetSetupManager()
	params := setupManager.GetParams()
	log.Printf("params = %+v", params)

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Params"] = params
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
