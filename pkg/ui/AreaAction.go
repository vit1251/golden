package ui

import (
	"net/http"
	"path/filepath"
	"html/template"
	"fmt"
)

type AreaAction struct {
	Action
	tmpl     *template.Template
}

func NewAreaAction() (*AreaAction) {
	aa := new(AreaAction)

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo_index.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	aa.tmpl = tmpl

	return aa
}

func (self *AreaAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse parameters */
	webSite := self.Site

	/* Get area manager */
	areaManager := webSite.GetAreaManager()

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Rener */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areas
	self.tmpl.ExecuteTemplate(w, "layout", outParams)

}
