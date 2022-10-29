package action

import (
	"fmt"
	"github.com/vit1251/golden/internal/i18n"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
	"strings"
)

type FileEchoIndexAction struct {
	Action
}

func NewFileEchoIndexAction() *FileEchoIndexAction {
	aa := new(FileEchoIndexAction)
	return aa
}

func (self *FileEchoIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	//fileMapper := mapperManager.GetFileMapper()
	fileAreaMapper := mapperManager.GetFileAreaMapper()

	/* Get message area */
	simpleAreas, err1 := fileAreaMapper.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail in GetAreas on fileAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	areasWithCounter, err2 := fileAreaMapper.UpdateFileAreasWithFileCount(simpleAreas)
	if err2 != nil {
		response := fmt.Sprintf("Fail in UpdateFileAreasWithFileCount on fileAreaMapper: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	areasWithNewCounter, err3 := fileAreaMapper.UpdateNewFileAreasWithFileCount(areasWithCounter)
	if err3 != nil {
		response := fmt.Sprintf("Fail in UpdateNewFileAreasWithFileCount on fileAreaMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	areas := areasWithNewCounter

	/* Render */
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
		SetClass("file-index-table")

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

func (self *FileEchoIndexAction) renderRow(area *mapper.FileArea) widgets2.IWidget {

	rowTitle := fmt.Sprintf("[%d] %s - %s (%s)",
		area.GetOrder(), area.GetName(), area.GetSummary(), area.GetCharset(),
	)

	/* Make message row container */
	rowWidget := widgets2.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center").
		SetTitle(rowTitle)

	var classNames []string
	classNames = append(classNames, "echo-index-item")
	//if area.NewMessageCount > 0 {
	//	classNames = append(classNames, "echo-index-item-new")
	//}
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render area name */
	nameWidget := widgets2.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(area.GetName())
	rowWidget.AddWidget(nameWidget)

	/* Render summary */
	summaryWidget := widgets2.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
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
		SetWidth("180px").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		SetStyle("flex-shrink: 0").
		//SetStyle("border: 1px solid blue").
		AddWidget(counterWidgetContent)
	rowWidget.AddWidget(counterWidget)

	/* Link container */
	navigateAddress := fmt.Sprintf("/file/%s", area.GetName())

	navigateItem := widgets2.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}

func (self *FileEchoIndexAction) renderMessageCounter(area *mapper.FileArea) widgets2.IWidget {
	counterWidget := widgets2.NewDivWidget()

	newCount := area.GetNewCount()
	if newCount > 0 {
		msgCount := fmt.Sprintf("%d", newCount)
		counterWidget.SetContent(msgCount)
	}

	return counterWidget
}

func (self *FileEchoIndexAction) renderActions() widgets2.IWidget {

	var mainLanguage string = i18n.GetDefaultLanguage()

	/* Render action bar */
	actionBar := widgets2.NewActionMenuWidget()

	actionLabel := i18n.GetText(mainLanguage, "FileEchoIndexAction", "action-button-create")
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/file/create")).
		SetIcon("icofont-update").
		SetLabel(actionLabel))

	return actionBar

}
