package ui

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type NetmailComposeAction struct {
	Action
	tmpl     *template.Template
}

func NewNetmailComposeAction() (*NetmailComposeAction) {
	nm := new(NetmailComposeAction)

	/* Cache HTML page template */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "netmail_compose.tmpl")
	tmpl, err1 := template.ParseFiles(lp, fp)
	if err1 != nil {
		panic(err1)
	}
	nm.tmpl = tmpl

	return nm
}

func (self *NetmailComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	outParams := make(map[string]interface{})
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
