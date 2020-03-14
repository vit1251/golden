package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ui/views"
	version2 "github.com/vit1251/golden/pkg/version"
	"net/http"
	"path/filepath"
)

type WelcomeAction struct {
	Action
}

func NewWelcomeAction() *WelcomeAction {
	wa := new(WelcomeAction)
	return wa
}

func (self *WelcomeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Get dependency injection manager */
	version := version2.GetVersion()

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "welcome.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Version", version)
	if err := doc.Render(w); err != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
