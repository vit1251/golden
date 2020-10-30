package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type NetmailViewAction struct {
	Action
}

func NewNetmailViewAction() *NetmailViewAction {
	va := new(NetmailViewAction)
	return va
}

func (self *NetmailViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	netmailManager := self.restoreNetmailManager()
	//configManager := self.restoreConfigManager()

	//
	vars := mux.Vars(r)

	//
	msgHash := vars["msgid"]
	origMsg, err3 := netmailManager.GetMessageByHash(msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if origMsg == nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	content := origMsg.GetContent()

	//
	mtp := msg.NewMessageTextProcessor()
	err4 := mtp.Prepare(content)
	if err4 != nil {
		response := fmt.Sprintf("Fail on Prepare on MessageTextProcessor")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	outDoc := mtp.HTML()

	/* Update view counter */
	err5 := netmailManager.ViewMessageByHash(msgHash)
	if err5 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash on messageManager: err = %+v", err5)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/netmail/%s/reply", origMsg.Hash)).
			SetClass("netmail-reply-action").
			SetIcon("icofont-edit").
			SetLabel("Reply")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/netmail/%s/remove", origMsg.Hash)).
			SetClass("netmail-remove-action").
			SetIcon("icofont-remove").
			SetLabel("Delete"))
	containerVBox.Add(amw)

	indexTable := widgets.NewTableWidget().
		SetClass("table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("FROM"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(origMsg.From))))

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("TO"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(origMsg.To))))

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("SUBJ"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(origMsg.Subject))))

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("DATE"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(
			fmt.Sprintf("%s", origMsg.DateWritten)))))

	containerVBox.Add(indexTable)

	previewWidget := widgets.NewDivWidget().
		SetClass("message-preview").
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
