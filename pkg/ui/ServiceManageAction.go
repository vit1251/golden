package ui

import (
	"github.com/vit1251/golden/pkg/common"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type ServiceManageAction struct {
	Action
	tmpl *template.Template
}

func NewServiceManageAction() *ServiceManageAction {
	sma := new(ServiceManageAction)

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "service_index.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	sma.tmpl = tmpl

	return sma
}

func (self *ServiceManageAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

	/* Setup manager operation */
	setupManager := master.SetupManager
	params := setupManager.GetParams()
	log.Printf("params = %+v", params)

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Params"] = params
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}

