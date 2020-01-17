package ui

import (
	"net/http"
//	"github.com/gorilla/mux"
//	msgProc "github.com/vit1251/golden/pkg/msg"
	"path/filepath"
	"html/template"
//	"log"
)

type StatAction struct {
	Action
	tmpl     *template.Template   /* Page template cache   */
}

func NewStatAction() (*StatAction) {

	/* New statistics action */
	sa := new(StatAction)

	/* Cache HTML page template */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "stat.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	sa.tmpl = tmpl

	return sa
}

func (self *StatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Render */
	outParams := make(map[string]interface{})
	self.tmpl.ExecuteTemplate(w, "layout", outParams)

}
