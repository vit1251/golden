package ui

import (
	"fmt"
	stat2 "github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/ui/views"
	"net/http"
	"path/filepath"
)

type StatAction struct {
	Action
}

func NewStatAction() *StatAction {
	sa := new(StatAction)
	return sa
}

func (self *StatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var statManager *stat2.StatManager
	self.Container.Invoke(func(sm *stat2.StatManager) {
		statManager = sm
	})

	/* Get statistics */
	stat, err1 := statManager.GetStat()
	if err1 != nil {
		response := fmt.Sprintf("Fail GetStat on StatManager: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create statistics */
	Status := make(map[string]string)

	Status["Total TIC Received"] = fmt.Sprintf("%d", stat.TicReceived)
	Status["Total TIC Sent"] = fmt.Sprintf("%d", stat.TicSent)

	Status["Total Echomail Received"] = fmt.Sprintf("%d", stat.EchomailReceived)
	Status["Total Echomail Sent"] = fmt.Sprintf("%d", stat.EchomailSent)

	Status["Total Packet Received"] = fmt.Sprintf("%d", stat.PacketReceived)
	Status["Total Packet Sent"] = fmt.Sprintf("%d", stat.PacketSent)

	Status["Total Message Received"] = fmt.Sprintf("%d", stat.MessageReceived)
	Status["Total Message Sent"] = fmt.Sprintf("%d", stat.MessageSent)

	Status["Dupe Count"] = fmt.Sprintf("%d", stat.Dupe)

//	Status["Total Received"] = "N/A"
//	Status["Total Sent"] = "N/A"

//	Status["Total Ses In"] = "N/A"
//	Status["Total Ses Out"] = "N/A"

//	Status["Total Sessions"] = "N/A"

//	Status["Time In"] = "N/A"
//	Status["Time Out"] = "N/A"

//	Status["Time Online"] = "N/A"

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "stat_index.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Status", Status)
	if err := doc.Render(w); err != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
}
