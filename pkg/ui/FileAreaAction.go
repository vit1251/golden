package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/common"
	"html/template"
	"net/http"
	"path/filepath"
)

type FileAreaAction struct {
	Action
	tmpl     *template.Template
}

func NewFileAreaAction() *FileAreaAction {
	aa := new(FileAreaAction)

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "file_area_index.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	aa.tmpl = tmpl

	return aa
}

func (self *FileAreaAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()
	fileAreaManager := master.FileManager

	/* Get message area */
	areas, err1 := fileAreaManager.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areas
	self.tmpl.ExecuteTemplate(w, "layout", outParams)

}
