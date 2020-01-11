package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
	"fmt"
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
		response := fmt.Sprintf("Fail on ParseFiles")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Parse parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	webSite := self.Site

	/* Get area manager */
	areaManager := webSite.GetAreaManager()
	area, err1 := areaManager.GetAreaByName(echoTag)
	if (err1 != nil) {
		response := fmt.Sprintf("Fail on GetAreaByName where echoTag is %s: err = %+v", echoTag, err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %v", area)

	/* Get message headers */
	messageManager := webSite.GetMessageManager()
	msgHeaders, err2 := messageManager.GetMessageHeaders(echoTag)
	if (err2 != nil) {
		response := fmt.Sprintf("Fail on GetMessageHeaders where echoTag is %s: err = %+v", echoTag, err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
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
