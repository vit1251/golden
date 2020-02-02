package ui

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type ServiceManageAction struct {
	Action
}

func NewServiceManageAction() *ServiceManageAction {
	sma := new(ServiceManageAction)
	return sma
}

func (self *ServiceManageAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "service_index.tmpl")
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

