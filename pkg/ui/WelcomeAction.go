package ui

import (
	version2 "github.com/vit1251/golden/pkg/version"
	"net/http"
	"html/template"
	//	"github.com/gorilla/mux"
	"path/filepath"
)

type WelcomeAction struct {
	Action
}

func NewWelcomeAction() (*WelcomeAction) {
	wa := new(WelcomeAction)
	return wa
}

func (self *WelcomeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "welcome.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

	/* Get dependency injection manager */
	version := version2.GetVersion()

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Version"] = version
	tmpl.ExecuteTemplate(w, "layout", outParams)

}
