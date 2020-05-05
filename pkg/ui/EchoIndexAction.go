package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoIndexAction struct {
	Action
}

func NewEchoIndexAction() *EchoIndexAction {
	aa := new(EchoIndexAction)
	return aa
}

func (self *EchoIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *msg.AreaManager
	self.Container.Invoke(func(am *msg.AreaManager) {
		areaManager = am
	})

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	indexTable := widgets.NewTableWidget().
		SetClass("echo-index-items")

	indexTable.
		AddRow(widgets.NewTableRowWidget().
			SetClass("echo-index-header").
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Summary"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Count"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, area := range areas {
		log.Printf("area = %+v", area)
		row := widgets.NewTableRowWidget()

		if area.NewMessageCount > 0 {
			row.SetClass("echo-index-item-new")
		} else {
			row.SetClass("echo-index-item")
		}

		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(area.Name())))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(area.Summary)))

		if area.NewMessageCount > 0 {

			hBox := widgets.NewVBoxWidget()

			newMsgCount := widgets.NewTextWidgetWithText(fmt.Sprintf("%d", area.NewMessageCount))
			separator := widgets.NewTextWidgetWithText(" / ")
			msgCount :=	widgets.NewTextWidgetWithText(fmt.Sprintf("%d", area.MessageCount))

			newMsgCount.SetClass("echo-index-item-count-new")

			hBox.Add(newMsgCount)
			hBox.Add(separator)
			hBox.Add(msgCount)

			cell := widgets.NewTableCellWidget().SetWidget(hBox)
			cell.SetClass("echo-index-item-count")
			row.AddCell(cell)

		} else {
			cell := widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(fmt.Sprintf("%d", area.MessageCount)))
			cell.SetClass("echo-index-item-count")
			row.AddCell(cell)
		}

		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("View").
				SetClass("btn").
				SetLink(fmt.Sprintf("/echo/%s", area.Name()))))


		indexTable.AddRow(row)
	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}


}
