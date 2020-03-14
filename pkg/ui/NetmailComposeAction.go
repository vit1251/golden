package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ui/views"
	"net/http"
	"path/filepath"
)

type NetmailComposeAction struct {
	Action
}

func NewNetmailComposeAction() (*NetmailComposeAction) {
	nm := new(NetmailComposeAction)
	return nm
}

func (self *NetmailComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "netmail_compose.tmpl")
	doc.SetPage(pagePath)
	err1 := doc.Render(w)
	if err1 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
