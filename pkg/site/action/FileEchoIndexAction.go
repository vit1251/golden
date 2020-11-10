package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type FileEchoIndexAction struct {
	Action
}

func NewFileEchoIndexAction() *FileEchoIndexAction {
	aa := new(FileEchoIndexAction)
	return aa
}

func (self *FileEchoIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	//fileMapper := mapperManager.GetFileMapper()
	fileAreaMapper := mapperManager.GetFileAreaMapper()

	/* Get message area */
	simpleAreas, err1 := fileAreaMapper.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail in GetAreas on fileAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	areas, err2 := fileAreaMapper.UpdateFileAreasWithFileCount(simpleAreas)
	if err2 != nil {
		response := fmt.Sprintf("Fail in UpdateFileAreasWithFileCount on fileAreaMapper: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

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

	indexTable := widgets.NewTableWidget().
		SetClass("file-index-table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		SetClass("file-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Summary"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Count"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, area := range areas {
		log.Printf("area = %+v", area)

		actions := widgets.NewVBoxWidget()

		actions.Add(widgets.NewLinkWidget().
			SetContent("View").
			SetClass("btn").
			SetLink(fmt.Sprintf("/file/%s", area.GetName())))

		indexTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(area.GetName()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(area.Summary))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(
				fmt.Sprintf("%d", area.Count)))).
			AddCell(widgets.NewTableCellWidget().SetWidget(actions)))
	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
