package ui

import (
	"github.com/vit1251/golden/pkg/common"
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
	fp := filepath.Join("views", "netmail_index.tmpl")
	tmpl, err1 := template.ParseFiles(lp, fp)
	if err1 != nil {
		panic(err1)
	}
	nm.tmpl = tmpl

	return nm
}

func (self *NetmailAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()
	msgHeaders, err1 := master.NetmailManager.GetMessageHeaders()
	if err1 != nil {
		panic(err1)
	}

	var actions []*UserAction
	action1 := NewUserAction()
	action1.Link = "/netmail/compose"
	action1.Icon = "/static/img/icon/quote-50.png"
	action1.Label = "Compose"
	actions = append(actions, action1)

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Actions"] = actions
	outParams["msgHeaders"] = msgHeaders
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
