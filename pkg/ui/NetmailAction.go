package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/ui/views"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"net/http"
	"path/filepath"
)

type NetmailAction struct {
	Action
}

func NewNetmailAction() (*NetmailAction) {
	nm := new(NetmailAction)
	return nm
}

func (self *NetmailAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var netmailManager *netmail.NetmailManager
	self.Container.Invoke(func(nm *netmail.NetmailManager) {
		netmailManager = nm
	})

	/* Message headers */
	msgHeaders, err1 := netmailManager.GetMessageHeaders()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders on NetmailManager: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create menu actions */
	mw := widgets.NewMenuWidget()
	action1 := widgets.NewMenuAction()
	action1.Link = "/netmail/compose"
	action1.Icon = "/static/img/icon/quote-50.png"
	action1.Label = "Compose"
	mw.Add(action1)

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "netmail_index.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Actions", mw.Actions())
	doc.SetParam("msgHeaders", msgHeaders)
	err3 := doc.Render(w)
	if err3 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
