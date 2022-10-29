package action

import (
	"fmt"
	"github.com/vit1251/golden/internal/i18n"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/internal/utils"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
	"strings"
)

type EchoAreaIndexAction struct {
	Action
}

func NewEchoAreaIndexAction() *EchoAreaIndexAction {
	aa := new(EchoAreaIndexAction)
	return aa
}

func (self *EchoAreaIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Get message area */
	areas, err1 := echoAreaMapper.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Update area index default identification */
	for _, area := range areas {
		var areaIndex string = area.GetAreaIndex()
		if areaIndex == "" {
			var areaIndex string = utils.IndexHelper_makeUUID()
			area.SetAreaIndex(areaIndex)
			echoAreaMapper.Update(&area)
		}
	}

	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	/* Context actions */
	actionsBar := self.renderActions()
	containerVBox.Add(actionsBar)

	indexTable := widgets2.NewDivWidget().
		SetClass("echo-index-table")

	for _, area := range areas {

		areaRow := self.renderRow(&area)
		indexTable.AddWidget(areaRow)

	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *EchoAreaIndexAction) renderMessageCounter(area *mapper.Area) widgets2.IWidget {

	counterWidget := widgets2.NewDivWidget()

	if area.GetNewMessageCount() > 0 {

		newMsgCount := widgets2.NewTextWidgetWithText(fmt.Sprintf("%d", area.GetNewMessageCount()))
		newMsgCount.SetClass("echo-index-item-count-new")

		counterWidget.AddWidget(newMsgCount)
		counterWidget.SetClass("echo-index-item-count")

	}

	return counterWidget

}

func (self *EchoAreaIndexAction) renderRow(area *mapper.Area) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())

	rowTitle := fmt.Sprintf("[%d] %s - %s (%s)",
		area.GetOrder(), area.GetName(), area.GetSummary(), area.GetCharset(),
	)

	/* Make message row container */
	rowWidget := widgets2.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("flex-direction: row").
		SetStyle("align-items: center").
		SetTitle(rowTitle)

	var classNames []string
	classNames = append(classNames, "echo-index-item")
	if area.GetNewMessageCount() > 0 {
		classNames = append(classNames, "echo-index-item-new")
	}
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render area name */
	nameWidget := widgets2.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: flex-start").
		SetStyle("justify-content: center").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(area.GetName())
	rowWidget.AddWidget(nameWidget)

	/* Render NEW point */
	var newPointContent string = ""
	if area.GetNewMessageCount() > 0 {
		newPointContent = "•"
	}
	newPointWidget := widgets2.NewDivWidget().
		SetWidth("20px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: center").
		SetStyle("justify-content: center").
		//SetStyle("border: 1px solid green").
		SetStyle("color: yellow").
		SetContent(newPointContent)
	rowWidget.AddWidget(newPointWidget)

	/* Render summary */
	summaryWidget := widgets2.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: flex-start").
		SetStyle("justify-content: center").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid red").
		SetContent(area.GetSummary())
	rowWidget.AddWidget(summaryWidget)

	/* Render counter widget */
	counterWidgetContent := self.renderMessageCounter(area)
	counterWidget := widgets2.NewDivWidget().
		SetHeight("38px").
		SetWidth("160px").
		SetStyle("flex-shrink: 0").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: flex-end").
		SetStyle("justify-content: center").
		AddWidget(counterWidgetContent)
	rowWidget.AddWidget(counterWidget)

	/* Link container */

	navigateAddress := urlManager.CreateUrl("/echo/{area_index}").
		SetParam("area_index", area.GetAreaIndex()).
		Build()

	navigateItem := widgets2.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}

func (self *EchoAreaIndexAction) renderActions() widgets2.IWidget {

	var mainLanguage string = i18n.GetDefaultLanguage()

	/* Render action bar */
	actionBar := widgets2.NewActionMenuWidget()

	actionLabel := i18n.GetText(mainLanguage, "EchoAreaIndexAction", "action-button-create")
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/echo/create")).
		SetIcon("icofont-update").
		SetLabel(actionLabel))

	return actionBar

}
