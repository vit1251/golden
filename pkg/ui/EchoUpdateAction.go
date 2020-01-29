package ui

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type EchoUpdateAction struct {
	Action
	tmpl  *template.Template
}

func NewEchoUpdateAction() (*EchoAction) {
	ea := new(EchoAction)

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
	outParams := make(map[string]interface{})
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
