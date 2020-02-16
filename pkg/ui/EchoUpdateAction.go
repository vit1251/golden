package ui

import (
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/common"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type EchoUpdateAction struct {
	Action
	tmpl  *template.Template
}

func NewEchoUpdateAction() *EchoUpdateAction {
	ea := new(EchoUpdateAction)

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_update.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	ea.tmpl = tmpl

	return ea
}

func (self *EchoUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	areaManager := master.AreaManager
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %v", area)

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Area"] = area
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
