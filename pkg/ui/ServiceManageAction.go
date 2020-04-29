package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
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

	/* Calculate summary */
	var newDirectMsgCount int
	var newEchoMsgCount int
	var newFileCount int
	self.Container.Invoke(func(nm *netmail.NetmailManager, em *msg.MessageManager, fm *file.FileManager) {
		newDirectMsgCount, _ = nm.GetMessageNewCount()
		newEchoMsgCount, _ = em.GetMessageNewCount()
		newFileCount, _ = fm.GetMessageNewCount()
	})

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	mmw.SetParam("mainMenuDirect", newDirectMsgCount)
	mmw.SetParam("mainMenuEcho", newEchoMsgCount)
	mmw.SetParam("mainMenuFile", newFileCount)
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

