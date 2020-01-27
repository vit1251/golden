package ui

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type NetmailAction struct {
	Action
	tmpl     *template.Template
}

func NewNetmailAction() (*NetmailAction) {
	nm := new(NetmailAction)

	/* Cache HTML page template */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "netmail.tmpl")
	tmpl, err1 := template.ParseFiles(lp, fp)
	if err1 != nil {
		panic(err1)
	}
	nm.tmpl = tmpl

	return nm
}

func (self *NetmailAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	outParams := make(map[string]interface{})
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
