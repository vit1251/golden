package ui

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type FileAction struct {
	Action
	tmpl  *template.Template
}

func NewFileAction() (*FileAction) {
	fa := new(FileAction)

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "file.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	fa.tmpl = tmpl
	return fa
}


func (self *FileAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Not implemented yet")
	http.Error(w, response, http.StatusInternalServerError)
}
