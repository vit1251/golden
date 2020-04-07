package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"net/http"
)

type HelpAction struct {
	Action
}

func NewHelpAction() *HelpAction {
	wa := new(HelpAction)
	return wa
}

func (self *HelpAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	content := self.makeHelpContent()

	documentWidget := widgets.NewDivWidget().
		SetContent(content).
		SetClass("pre-wrap")
	vBox.Add(documentWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *HelpAction) makeHelpContent() string {
	var result string

	result += "*User steps*\n"
	result += "1. Setup your parameters in page *Setup*\n"
	result += "2. Start mailer and tosser on page *Service*\n"
	result += "3. Write NetMail and EchoMail on page *Netmail* and *Echomail*\n"
	result += "4. Research file echos archive in *Filebox*\n"
	result += "5. Wash your hands and be happy ;)\n"

	return result
}

