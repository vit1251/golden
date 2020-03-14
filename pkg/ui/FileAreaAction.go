package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/ui/views"
	"net/http"
	"path/filepath"
)

type FileAreaAction struct {
	Action
}

func NewFileAreaAction() *FileAreaAction {
	aa := new(FileAreaAction)
	return aa
}

func (self *FileAreaAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var fileManager *file.FileManager
	self.Container.Invoke(func(fm *file.FileManager) {
		fileManager = fm
	})

	/* Get message area */
	areas, err1 := fileManager.GetAreas2()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "file_area_index.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Areas", areas)
	err2 := doc.Render(w)
	if err2 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}
