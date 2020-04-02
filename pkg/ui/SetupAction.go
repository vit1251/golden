package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/ui/views"
	"log"
	"net/http"
	"path/filepath"
)

type SetupAction struct {
	Action
}

func NewSetupAction() (*SetupAction) {
	sa := new(SetupAction)
	return sa
}

func (self *SetupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Setup manager operation */
	var setupManager *setup.ConfigManager
	self.Container.Invoke(func(sm *setup.ConfigManager) {
		setupManager = sm
	})

	params := setupManager.GetParams()
	log.Printf("params = %+v", params)

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "setup.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Params", params)
	if err := doc.Render(w); err != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
