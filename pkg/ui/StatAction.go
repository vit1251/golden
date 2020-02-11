package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/stat"
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
	fp := filepath.Join("views", "stat_index.tmpl")
	tmpl, err1 := template.ParseFiles(lp, fp)
	if err1 != nil {
		panic(err1)
	}
	sa.tmpl = tmpl

	return sa
}

func (self *StatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	statm := stat.NewStatManager()
	stat, err1 := statm.GetStat()
	if err1 != nil {
		panic(err1)
	}

	/* Create statistics */
	Status := make(map[string]string)

	Status["Total TIC Received"] = fmt.Sprintf("%d", stat.TicReceived)
	Status["Total TIC Sent"] = fmt.Sprintf("%d", stat.TicSent)

	Status["Total Echomail Received"] = fmt.Sprintf("%d", stat.EchomailReceived)
	Status["Total Echomail Sent"] = fmt.Sprintf("%d", stat.EchomailSent)

//	Status["Total Received"] = "N/A"
//	Status["Total Sent"] = "N/A"

//	Status["Total Ses In"] = "N/A"
//	Status["Total Ses Out"] = "N/A"

//	Status["Total Sessions"] = "N/A"

//	Status["Time In"] = "N/A"
//	Status["Time Out"] = "N/A"

//	Status["Time Online"] = "N/A"

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Status"] = Status
	self.tmpl.ExecuteTemplate(w, "layout", outParams)
}
