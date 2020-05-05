package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"net/http"
)

type NetmailComposeAction struct {
	Action
}

func NewNetmailComposeAction() (*NetmailComposeAction) {
	nm := new(NetmailComposeAction)
	return nm
}

func (self *NetmailComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets.NewDivWidget().SetClass("container")
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.SetWidget(containerVBox)

	section := widgets.NewSectionWidget().
		SetTitle("Create netmail message")

	composeForm := widgets.NewFormWidget().
		SetAction("/netmail/compose/complete").
		SetMethod("POST")

	composeForm.SetWidget(widgets.NewVBoxWidget().
		Add(widgets.NewFormInputWidget().SetTitle("ToName").SetName("to").SetPlaceholder("Vitold Sedyshev")).
		Add(widgets.NewFormInputWidget().SetTitle("ToAddr").SetName("to_addr").SetPlaceholder("2:5023/24.3752")).
		Add(widgets.NewFormInputWidget().SetTitle("Subject").SetName("subject").SetPlaceholder("RE: Hello, world!")).
		Add(widgets.NewFormTextWidget().SetName("body")).
		Add(widgets.NewFormButtonWidget().SetType("submit").SetTitle("Send")))

	section.SetWidget(composeForm)
	containerVBox.Add(section)

	bw.SetWidget(vBox)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
