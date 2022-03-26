package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/i18n"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
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

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Get message area */
	areas, err1 := echoAreaMapper.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	/* Context actions */
	actionsBar := self.renderActions()
	containerVBox.Add(actionsBar)

	indexTable := widgets.NewDivWidget().
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

func (self *EchoAreaIndexAction) renderMessageCounter(area *mapper.Area) widgets.IWidget {

	counterWidget := widgets.NewDivWidget()

	if area.GetNewMessageCount() > 0 {

		newMsgCount := widgets.NewTextWidgetWithText(fmt.Sprintf("%d", area.GetNewMessageCount()))
		newMsgCount.SetClass("echo-index-item-count-new")

		counterWidget.AddWidget(newMsgCount)
		counterWidget.SetClass("echo-index-item-count")

	}

	return counterWidget

}

func (self *EchoAreaIndexAction) renderRow(area *mapper.Area) widgets.IWidget {

	rowTitle := fmt.Sprintf("[%d] %s - %s (%s)",
		area.GetOrder(), area.GetName(), area.GetSummary(), area.GetCharset(),
	)

	/* Make message row container */
	rowWidget := widgets.NewDivWidget().
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
	nameWidget := widgets.NewDivWidget().
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
	newPointWidget := widgets.NewDivWidget().
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
	summaryWidget := widgets.NewDivWidget().
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
	counterWidget := widgets.NewDivWidget().
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
	navigateAddress := fmt.Sprintf("/echo/%s", area.GetName())

	navigateItem := widgets.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}

func (self *EchoAreaIndexAction) renderActions() widgets.IWidget {

	var mainLanguage string = i18n.GetDefaultLanguage()

	/* Render action bar */
	actionBar := widgets.NewActionMenuWidget()

	actionLabel := i18n.GetText(mainLanguage, "EchoAreaIndexAction", "action-button-create")
	actionBar.Add(widgets.NewMenuAction().
		SetLink(fmt.Sprintf("/echo/create")).
		SetIcon("icofont-update").
		SetLabel(actionLabel))

	return actionBar

}
