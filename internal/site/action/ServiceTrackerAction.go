package action

import (
	"fmt"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"net/http"
)

type ServiceTrackerAction struct {
	Action
}

func NewServiceTrackerAction() *ServiceTrackerAction {
	sa := new(ServiceTrackerAction)
	return sa
}

func (self *ServiceTrackerAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Render */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	/* Context actions */
	amw := widgets2.NewActionMenuWidget().
		Add(widgets2.NewMenuAction().
			SetLink(fmt.Sprintf("/service/tracker/event")).
			SetIcon("icofont-update").
			SetLabel("Run"))

	containerVBox.Add(amw)

	container.AddWidget(containerVBox)

	vBox.Add(container)

	statWidget := widgets2.NewTableWidget().
		SetClass("stat-index-table")

	statWidget.AddRow(widgets2.NewTableRowWidget().
		SetClass("stat-index-header").
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Metric"))).
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Received"))).
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Sent"))))

	containerVBox.Add(statWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
