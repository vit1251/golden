package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/common"
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

	master := common.GetMaster()
	fileManager := master.FileManager

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
	outParams := make(map[string]interface{})
	outParams["Area"] = area
	outParams["Files"] = files
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
