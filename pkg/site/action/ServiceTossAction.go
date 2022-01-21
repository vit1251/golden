package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type ServiceTossAction struct {
	Action
}

func NewServiceTossAction() *ServiceTossAction {
	sa := new(ServiceTossAction)
	return sa
}

func (self *ServiceTossAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/service/toss/event")).
			SetIcon("icofont-update").
			SetLabel("Run"))

	containerVBox.Add(amw)

	container.AddWidget(containerVBox)

	vBox.Add(container)

	statWidget := widgets.NewTableWidget().
		SetClass("stat-index-table")

	statWidget.AddRow(widgets.NewTableRowWidget().
		SetClass("stat-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Metric"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Received"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Sent"))))

	containerVBox.Add(statWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
