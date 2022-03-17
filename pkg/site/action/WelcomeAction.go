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



	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	mainWidget := widgets.NewDivWidget().
		SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	mainWidget.AddWidget(containerVBox)

	vBox.Add(mainWidget)

	/* Golden Point mascot image */

	imageWidget := self.renderVerpic()
	containerVBox.Add(imageWidget)

	/* Golden Point version */
	productWidget := self.renderProductVersion()
	containerVBox.Add(productWidget)

	/* Contributors */
	contributorWidget := self.renderContributors()
	containerVBox.Add(contributorWidget)

	/* Source code */
	sourceWidget := self.renderSourceCode()
	containerVBox.Add(sourceWidget)

	/* Donation*/
	donateWidget := self.renderDonation()
	containerVBox.Add(donateWidget)

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *WelcomeAction) renderVerpic() widgets.IWidget {

//	var version string = "1_2_16"
//	var version string = "1_2_17"
	var version string = "1_2_18"

	imageName := fmt.Sprintf("Dog_%s.png", version)
	imagePath := fmt.Sprintf("/static/%s", imageName)

	imageWidget := widgets.NewImageWidget()
	imageWidget.SetSource(imagePath).
		SetClass("welcome-img")

	return imageWidget

}

func (self *WelcomeAction) renderProductVersion() widgets.IWidget {

	/* Get dependency injection manager */
	version := cmn.GetVersion()

	productWidget := widgets.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	/* Product name */
	nameWidget := widgets.NewDivWidget().
		SetClass("welcome-header").
		SetContent("Golden point").
		SetStyle("padding-bottom: 8px")

	productWidget.AddWidget(nameWidget)

	/* Product version */
	versionWidget := widgets.NewDivWidget().
		SetStyle("text-align: center").
		SetContent(fmt.Sprintf("Version %s", version))

	productWidget.AddWidget(versionWidget)

	return productWidget

}

func (self *WelcomeAction) renderContributors() widgets.IWidget {

	contributorWidget := widgets.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	contributorHeader := widgets.NewDivWidget().
		SetClass("welcome-contributor-header").
		SetContent("Contributors").
		SetStyle("padding-bottom: 8px")

	contributorWidget.AddWidget(contributorHeader)

	contributors := cmn.GetContributors()
	var newContributros []string
	for _, c := range contributors {
		newContributros = append(newContributros, c.Name)
	}
	newContrib := strings.Join(newContributros, ", ")

	contributorList := widgets.NewDivWidget().
		SetClass("welcome-contributor-list").
		SetStyle("text-align: center").
		SetContent(newContrib)

	contributorWidget.AddWidget(contributorList)

	return contributorWidget

}

func (self *WelcomeAction) renderSourceCode() widgets.IWidget {

	sourceWidget := widgets.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	sourceHeaderWidget := widgets.NewDivWidget().
		SetClass("welcome-source").
		SetContent("Source code and support").
		SetStyle("padding-bottom: 8px")
	sourceWidget.AddWidget(sourceHeaderWidget)

	sourceLink := widgets.NewLinkWidget().
		SetLink("https://github.com/vit1251/golden").
		SetContent("https://github.com/vit1251/golden").
		SetClass("welcome-source-link")
	sourceWidget.AddWidget(sourceLink)

	return sourceWidget

}

func (self *WelcomeAction) renderDonation() widgets.IWidget {

	donationWidget := widgets.NewDivWidget().
		SetStyle("padding-bottom: 32px")

//	donationHeaderWidget := widgets.NewDivWidget().
//		SetClass("welcome-donate").
//		SetContent("Donation").
//		SetStyle("padding-bottom: 8px")
//	donationWidget.AddWidget(donationHeaderWidget)

//	serviceList := widgets.NewDivWidget().
//		SetClass("welcome-donate-list").
//		SetStyle("text-align: center").
//		AddWidget(patreonLink)

//	donationWidget.AddWidget(serviceList)

	return donationWidget

}
