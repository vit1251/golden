package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type FileAreaIndexAction struct {
	Action
}

func NewFileAreaIndexAction() *FileAreaIndexAction {
	aa := new(FileAreaIndexAction)
	return aa
}

func (self *FileAreaIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var fileManager *file.FileManager
	self.Container.Invoke(func(fm *file.FileManager) {
		fileManager = fm
	})

	/* Get message area */
	areas, err1 := fileManager.GetAreas2()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	indexTable := widgets.NewTableWidget().
		SetClass("table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Summary"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Count"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, area := range areas {
		log.Printf("area = %+v", area)

		indexTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(area.Name()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(area.Summary))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(
				fmt.Sprintf("%d", area.Count)))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("View").
				SetLink(fmt.Sprintf("/file/%s", area.Name())))))
	}

	vBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
