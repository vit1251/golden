package ui

import (
	"github.com/vit1251/golden/pkg/common"
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

	master := common.GetMaster()

	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "setup.tmpl")
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
