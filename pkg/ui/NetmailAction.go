package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type NetmailAction struct {
	Action
}

func NewNetmailAction() (*NetmailAction) {
	nm := new(NetmailAction)
	return nm
}

func (self *NetmailAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Calculate summary */
	var newDirectMsgCount int
	var newEchoMsgCount int
	var newFileCount int
	self.Container.Invoke(func(nm *netmail.NetmailManager, em *msg.MessageManager, fm *file.FileManager) {
		newDirectMsgCount, _ = nm.GetMessageNewCount()
		newEchoMsgCount, _ = em.GetMessageNewCount()
		newFileCount, _ = fm.GetMessageNewCount()
	})

	var netmailManager *netmail.NetmailManager
	self.Container.Invoke(func(nm *netmail.NetmailManager) {
		netmailManager = nm
	})

	/* Message headers */
	msgHeaders, err1 := netmailManager.GetMessageHeaders()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders on NetmailManager: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeader = %+v", msgHeaders)

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	mmw.SetParam("mainMenuDirect", newDirectMsgCount)
	mmw.SetParam("mainMenuEcho", newEchoMsgCount)
	mmw.SetParam("mainMenuFile", newFileCount)
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
		SetClass("direct-index-items")

	indexTable.
		AddRow(widgets.NewTableRowWidget().
			SetClass("direct-index-header").
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("From"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("To"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Subject"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Date"))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, msg := range msgHeaders {
		log.Printf("msg = %+v", msg)

		row := widgets.NewTableRowWidget()

		if msg.ViewCount == 0 {
			row.SetClass("message-item-new")
		}

		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.From)))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.To)))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.Subject)))
		row.AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.Age())))

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
