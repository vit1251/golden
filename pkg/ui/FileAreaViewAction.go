package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/ui/views"
	"log"
	"net/http"
	"path/filepath"
)

type FileAreaViewAction struct {
	Action
}

func NewFileAreaViewAction() *FileAreaViewAction {
	fa := new(FileAreaViewAction)
	return fa
}


func (self *FileAreaViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var fileManager *file.FileManager
	self.Container.Invoke(func(fm *file.FileManager) {
		fileManager = fm
	})

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get message area */
	area, err1 := fileManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on FileManager")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	files, err2 := fileManager.GetFileHeaders(echoTag)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetFileHeaders on FileManager")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("files = %+v", files)

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "file_area_view.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Area", area)
	doc.SetParam("Files", files)
	err3 := doc.Render(w)
	if err3 != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
}
