package ui

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type FileAreaViewAction struct {
	Action
	tmpl  *template.Template
}

func NewFileAreaViewAction() (*FileAreaViewAction) {
	fa := new(FileAreaViewAction)

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "file_area_view.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	fa.tmpl = tmpl
	return fa
}


func (self *FileAreaViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	webSite := self.Site

	/* Get area manager */
	fileManager := webSite.GetFileAreaManager()

	files, err1 := fileManager.GetFileHeaders(echoTag)
	if err1 != nil {
		// TODO - process error
	}
	log.Printf("files = %+v", files)

	/* Rener */
	outParams := make(map[string]interface{})
	outParams["Files"] = files
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
