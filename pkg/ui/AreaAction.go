package ui

import (
	"net/http"
	"path/filepath"
	"html/template"
	"fmt"
)

type AreaAction struct {
	Action
}

func NewAreaAction() (*AreaAction) {
	return new(AreaAction)
}

func (self *AreaAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "area.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		response := fmt.Sprintf("Fail on ParseFiles")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Parse parameters */
	webSite := self.Site

	/* Get area manager */
	areaManager := webSite.GetAreaManager()

	/* Rener */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areaManager.GetAreas()
	tmpl.ExecuteTemplate(w, "layout", outParams)

}
