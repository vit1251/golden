package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type ServiceAction struct {
	Action
}

func NewServiceAction() *ServiceAction {
	return new(ServiceAction)
}

func (self ServiceAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	serviceTable := widgets.NewTableWidget().
		SetClass("table")

	/* Show header */
	serviceTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Service"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	/* Render mailer service */
	serviceTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Mailer"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
			SetContent("Event").
			SetClass("btn").
			SetLink(fmt.Sprintf("/service/%s/event", "mailer")))))

	/* Render mailer service */
	serviceTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Tosser"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
			SetContent("Event").
			SetClass("btn").
			SetLink(fmt.Sprintf("/service/%s/event", "tosser")))))

	/* Render mailer service */
	//serviceTable.AddRow(widgets.NewTableRowWidget().
	//	AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Ticker"))).
	//	AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
	//		SetContent("Event").
	//		SetClass("btn").
	//		SetLink(fmt.Sprintf("/service/%s/event", "ticker")))))

	containerVBox.Add(serviceTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
