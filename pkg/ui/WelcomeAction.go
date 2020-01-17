package ui

import (
	"net/http"
//	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
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
	webSite := self.Site
	version := webSite.GetVersion()

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Version"] = version
	tmpl.ExecuteTemplate(w, "layout", outParams)

}
