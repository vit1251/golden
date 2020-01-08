package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"log"
)

type EchoAction struct {
	Action
}

func NewEchoAction() (*EchoAction) {
	return new(EchoAction)
}

func (self *EchoAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

	/* Parse parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get area manager */
	areaManager := self.Site.app.GetAreaManager()
	area, err1 := areaManager.GetAreaByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)

	/* Get message headers */
	msgHeaders, err2 := self.Site.app.MessageBaseReader.GetMessageHeaders(echoTag)
	if (err2 != nil) {
		panic(err2)
	}
	log.Printf("msgHeaders = %q", msgHeaders)
	for _, msg := range msgHeaders {
		log.Printf("msg = %q", msg)
	}

	/* Rener */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areaManager.GetAreas()
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	tmpl.ExecuteTemplate(w, "layout", outParams)

}
