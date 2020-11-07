package action

import (
	"fmt"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
	"strings"
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
	version := cmn.GetVersion()

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	imageWidget := widgets.NewImageWidget()
	imageWidget.SetSource("/static/fido.svg").SetClass("welcome-img")
	containerVBox.Add(imageWidget)

	nameWidget := widgets.NewDivWidget().
		SetClass("welcome-header").
		SetContent("Golden point")
	containerVBox.Add(nameWidget)

	versionWidget := widgets.NewDivWidget()
	versionWidget.SetClass("welcome-version")
	versionWidget.SetContent(fmt.Sprintf("Version %s", version))
	containerVBox.Add(versionWidget)

	contributorHeader := widgets.NewDivWidget().
		SetClass("welcome-contributor-header").
		SetContent("Contributers")
	containerVBox.Add(contributorHeader)

	contributors := cmn.GetContributors()
	var newContributros []string
	for _, c := range contributors {
		newContributros = append(newContributros, c)
	}
	newContrib := strings.Join(newContributros, ", ")

	contributorList := widgets.NewDivWidget().
		SetClass("welcome-contributor-list").
		SetContent(newContrib)

	containerVBox.Add(contributorList)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
