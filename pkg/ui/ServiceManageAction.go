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
}

func NewServiceManageAction() *ServiceManageAction {
	sma := new(ServiceManageAction)
	return sma
}

func (self *ServiceManageAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "service_index.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

	/* Setup manager operation */
	setupManager := master.SetupManager
	params := setupManager.GetParams()
	log.Printf("params = %+v", params)

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Params"] = params
	tmpl.ExecuteTemplate(w, "layout", outParams)
}

