package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/ui/widgets"
	version2 "github.com/vit1251/golden/pkg/version"
	"net/http"
)

type WelcomeAction struct {
	Action
}

func NewWelcomeAction() *WelcomeAction {
	wa := new(WelcomeAction)
	return wa
}

func (self *WelcomeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Calculate summary */
	var newDirectMsgCount int
	var newEchoMsgCount int
	var newFileCount int
	self.Container.Invoke(func(nm *netmail.NetmailManager, em *msg.MessageManager, fm *file.FileManager) {
		newDirectMsgCount, _ = nm.GetMessageNewCount()
		newEchoMsgCount, _ = em.GetMessageNewCount()
		newFileCount, _ = fm.GetMessageNewCount()
	})

	/* Get dependency injection manager */
	version := version2.GetVersion()

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	mmw.SetParam("mainMenuDirect", newDirectMsgCount)
	mmw.SetParam("mainMenuEcho", newEchoMsgCount)
	mmw.SetParam("mainMenuFile", newFileCount)
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	imageWidget := widgets.NewImageWidget()
	imageWidget.SetSource("/static/img/fido.gif")
	containerVBox.Add(imageWidget)

	nameWidget := widgets.NewDivWidget().
		SetClass("welcomeHeader").
		SetContent("Golden point")
	containerVBox.Add(nameWidget)

	versionWidget := widgets.NewDivWidget()
	versionWidget.SetClass("welcomeVersion")
	versionWidget.SetContent(fmt.Sprintf("Version %s", version))
	containerVBox.Add(versionWidget)

	contributorHeader := widgets.NewDivWidget().
		SetClass("contributorHeader").
		SetContent("Contributers")
	containerVBox.Add(contributorHeader)

	contributorList := widgets.NewDivWidget().
		SetClass("contributorList").
		SetContent("Sergey Anohin")

	containerVBox.Add(contributorList)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
