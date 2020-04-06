package ui

import (
	"fmt"
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

	/* Get dependency injection manager */
	version := version2.GetVersion()

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	imageWidget := widgets.NewImageWidget()
	imageWidget.SetSource("/static/img/fido.gif")
	vBox.Add(imageWidget)

	nameWidget := widgets.NewDivWidget()
	nameWidget.SetClass("welcomeHeader")
	nameWidget.SetContent("Golden point")
	vBox.Add(nameWidget)

	versionWidget := widgets.NewDivWidget()
	versionWidget.SetClass("welcomeVersion")
	versionWidget.SetContent(fmt.Sprintf("Version %s", version))
	vBox.Add(versionWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
