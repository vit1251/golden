package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/ui/views"
	"net/http"
	"path/filepath"
)

type AreaAction struct {
	Action
}

func NewAreaAction() (*AreaAction) {
	aa := new(AreaAction)
	return aa
}

func (self *AreaAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area.AreaManager
	self.Container.Invoke(func(am *area.AreaManager) {
		areaManager = am
	})

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "echo_index.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Areas", areas)
	err2 := doc.Render(w)
	if err2 != nil {
		response := fmt.Sprintf("Fail on Render")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
}
