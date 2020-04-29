package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
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

	mmw := widgets.NewMainMenuWidget()
	mmw.SetParam("mainMenuDirect", newDirectMsgCount)
	mmw.SetParam("mainMenuEcho", newEchoMsgCount)
	mmw.SetParam("mainMenuFile", newFileCount)
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
