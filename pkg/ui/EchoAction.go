package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"path/filepath"
	"html/template"
	"log"
)

func (self *EchoAction) ServeHTTP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "echo.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	//
	echoTag := params.ByName("name")
	log.Printf("echoTag = %v", echoTag)
	//
	area, err1 := self.Site.app.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)
	//
	msgHeaders, err2 := self.Site.app.MessageBaseReader.GetMessageHeaders(echoTag)
	if (err2 != nil) {
		panic(err2)
	}
	log.Printf("msgHeaders = %q", msgHeaders)
	for _, msg := range msgHeaders {
		log.Printf("msg = %q", msg)
	}
	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.AreaList.Areas
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
