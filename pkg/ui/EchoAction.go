package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/vit1251/golden/pkg/msgapi"
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
	area, err1 := self.Site.app.config.AreaList.SearchByName(echoTag)
	if (err1 != nil) {
		panic(err1)
	}
	log.Printf("area = %v", area)
	//
	var msgBase = msgapi.SquishMessageBase{}
	msgHeaders, err2 := msgBase.ReadBase(area.Path)
	if (err2 != nil) {
		panic(err2)
	}
	//
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.config.AreaList.Areas
	outParams["Area"] = area
	outParams["Headers"] = msgHeaders
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
