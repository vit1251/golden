package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type NetmailIndexAction struct {
	Action
}

func NewNetmailIndexAction() *NetmailIndexAction {
	nm := new(NetmailIndexAction)
	return nm
}

func (self NetmailIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	netmailMapper := mapperManager.GetNetmailMapper()

	/* Message headers */
	msgHeaders, err1 := netmailMapper.GetMessageHeaders()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders on netmailMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeader = %+v", msgHeaders)

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink("/netmail/compose").
			SetIcon("icofont-edit").
			SetLabel("Compose"))
	containerVBox.Add(amw)

	container.SetWidget(containerVBox)
	vBox.Add(container)

	indexTable := widgets.NewTableWidget().
		SetClass("netmail-index-table")

	indexTable.
		AddRow(widgets.NewTableRowWidget().
			SetClass("netmail-index-header").
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("From"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("To"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Subject"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Date"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, msg := range msgHeaders {
		log.Printf("msg = %+v", msg)

		row := widgets.NewTableRowWidget()

		if msg.ViewCount == 0 {
			row.SetClass("netmail-index-item-new")
		}

		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.From)))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.To)))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.Subject)))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.GetAge())))

		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
			SetContent("View").
			SetClass("btn").
			SetLink(fmt.Sprintf("/netmail/%s/view", msg.Hash))))

		indexTable.AddRow(row)
	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
