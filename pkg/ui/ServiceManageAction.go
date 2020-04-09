package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"net/http"
)

type ServiceManageAction struct {
	Action
}

type ServiceInfo struct {
	Label string
	Class string
}

func NewServiceManageAction() *ServiceManageAction {
	sma := new(ServiceManageAction)
	return sma
}

func (self *ServiceManageAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	header := widgets.NewHeaderWidget()
	header.SetTitle("Manual manage service")
	vBox.Add(header)

	mainForm := widgets.NewFormWidget().
		SetAction("/service/complete").
		SetMethod("POST")

	mainFormContainer := widgets.NewVBoxWidget()
	mainForm.SetWidget(mainFormContainer)

	selectServiceWidget := widgets.NewFormSelectWidget().
		SetName("service").
		AddOption("Tosser Service", "tosser").
		AddOption("Mailer Service", "mailer")
	mainFormContainer.Add(selectServiceWidget)

	submitButton := widgets.NewFormButtonWidget().
		SetTitle("Start").
		SetType("submit")
	mainFormContainer.Add(submitButton)

	vBox.Add(mainForm)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

