package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
	"strings"
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

	var services []string
	services = append(services, "mailer")
	services = append(services, "tosser")
	services = append(services, "tracker")

	for _, s := range services {

		serviceName := strings.Title(s)

		/* Render mailer service */
		serviceTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(serviceName))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("Event").
				SetClass("btn").
				SetLink(fmt.Sprintf("/service/%s/event", s)))))

	}

	containerVBox.Add(serviceTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
