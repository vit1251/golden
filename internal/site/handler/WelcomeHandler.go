package handler

import (
	"fmt"
	"net/http"
	"strings"

	commonfunc "github.com/vit1251/golden/internal/common"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/registry"
)

type WelcomeHandler struct {
	registry *registry.Container
}

func NewWelcomeHandler(registry *registry.Container) *WelcomeHandler {
	return &WelcomeHandler{
		registry: registry,
	}
}

func (h *WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Render */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets2.NewMainMenuWidget()
	vBox.Add(mmw)

	mainWidget := widgets2.NewDivWidget().
		SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	mainWidget.AddWidget(containerVBox)

	vBox.Add(mainWidget)

	/* Golden Point mascot image */

	imageWidget := h.renderVerpic()
	containerVBox.Add(imageWidget)

	/* Golden Point version */
	productWidget := h.renderProductVersion()
	containerVBox.Add(productWidget)

	/* Community */
	donateWidget := h.renderCommunity()
	containerVBox.Add(donateWidget)

	/* Source code */
	sourceWidget := h.renderSourceCode()
	containerVBox.Add(sourceWidget)

	/* Contributors */
	contributorWidget := h.renderContributors()
	containerVBox.Add(contributorWidget)

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (h *WelcomeHandler) renderVerpic() widgets2.IWidget {

	//    var version string = "1_2_16"
	//    var version string = "1_2_17"
	//    var version string = "1_2_18"
	var version string = "1_2_19"

	imageName := fmt.Sprintf("Dog_%s.png", version)
	imagePath := fmt.Sprintf("/static/%s", imageName)

	imageWidget := widgets2.NewImageWidget()
	imageWidget.SetSource(imagePath).
		SetClass("welcome-img")

	return imageWidget

}

func (h *WelcomeHandler) renderProductVersion() widgets2.IWidget {

	/* Get dependency injection manager */
	version := commonfunc.GetVersion()

	productWidget := widgets2.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	/* Product name */
	nameWidget := widgets2.NewDivWidget().
		SetClass("welcome-header gold-text").
		SetContent("Golden Point").
		SetStyle("padding-bottom: 8px")

	productWidget.AddWidget(nameWidget)

	/* Product version */
	versionWidget := widgets2.NewDivWidget().
		SetStyle("text-align: center").
		SetContent(fmt.Sprintf("Version %s", version))

	productWidget.AddWidget(versionWidget)

	return productWidget

}

func (h *WelcomeHandler) renderContributors() widgets2.IWidget {

	contributorWidget := widgets2.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	contributorHeader := widgets2.NewDivWidget().
		SetClass("welcome-contributor-header").
		SetContent("Contributors").
		SetStyle("padding-bottom: 8px")

	contributorWidget.AddWidget(contributorHeader)

	contributors := commonfunc.GetContributors()
	var newContributros []string
	for _, c := range contributors {
		newContributros = append(newContributros, c.Name)
	}
	newContrib := strings.Join(newContributros, ", ")

	contributorList := widgets2.NewDivWidget().
		SetClass("welcome-contributor-list").
		SetStyle("text-align: center").
		SetContent(newContrib)

	contributorWidget.AddWidget(contributorList)

	return contributorWidget

}

func (h *WelcomeHandler) renderSourceCode() widgets2.IWidget {

	sourceWidget := widgets2.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	sourceHeaderWidget := widgets2.NewDivWidget().
		SetClass("welcome-source").
		SetContent("Source code and developing").
		SetStyle("padding-bottom: 8px")
	sourceWidget.AddWidget(sourceHeaderWidget)

	sourceLink := widgets2.NewLinkWidget().
		SetLink("https://github.com/vit1251/golden").
		SetContent("https://github.com/vit1251/golden").
		SetClass("welcome-source-link")
	sourceWidget.AddWidget(sourceLink)

	return sourceWidget

}

func (h *WelcomeHandler) renderCommunity() widgets2.IWidget {

	communityWidget := widgets2.NewDivWidget().
		SetStyle("padding-bottom: 32px")

	communityHeaderWidget := widgets2.NewDivWidget().
		SetClass("welcome-community").
		SetContent("User Group Community").
		SetStyle("padding-bottom: 8px")
	communityWidget.AddWidget(communityHeaderWidget)

	socialLink := widgets2.NewLinkWidget().
		SetLink("https://t.me/golden_point_community").
		SetContent("https://t.me/golden_point_community").
		SetClass("welcome-community-link")

	serviceList := widgets2.NewDivWidget().
		SetClass("welcome-community-list").
		SetStyle("text-align: center").
		AddWidget(socialLink)

	communityWidget.AddWidget(serviceList)

	return communityWidget

}
