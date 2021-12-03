package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/packet"
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

func (self NetmailViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	netmailMapper := mapperManager.GetNetmailMapper()

	//
	vars := mux.Vars(r)

	//
	msgHash := vars["msgid"]
	origMsg, err3 := netmailMapper.GetMessageByHash(msgHash)
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

	/* Preprocess message body (attachments) */
	rawPacket := origMsg.GetPacket()
	bodyParser := packet.NewMessageBodyParser()
	msgBody, _ := bodyParser.Parse(rawPacket)
	// TODO - use message parsing ... rawContent := msgBody.GetContent()
	// TODO - use message parsing ... content := string(rawContent)

	/* Processing message body */
	mtp := msg.NewMessageTextProcessor()
	err4 := mtp.Prepare(content)
	if err4 != nil {
		response := fmt.Sprintf("Fail on Prepare on MessageTextProcessor")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	outDoc := mtp.HTML()

	/* Update view counter */
	err5 := netmailMapper.ViewMessageByHash(msgHash)
	if err5 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash on netmailMapper: err = %+v", err5)
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

	container.AddWidget(containerVBox)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/netmail/%s/reply", origMsg.Hash)).
			SetClass("netmail-reply-action").
			SetIcon("icofont-edit").
			SetLabel("Reply")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/netmail/%s/dump", origMsg.Hash)).
			SetClass("netmail-dump-action").
			SetIcon("icofont-dump").
			SetLabel("Info")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/netmail/%s/remove", origMsg.Hash)).
			SetClass("netmail-remove-action").
			SetIcon("icofont-remove").
			SetLabel("Delete"))
	containerVBox.Add(amw)

	msgHeader := self.makeMessageHeaderSection(origMsg, msgBody)
	msgHeaderWrapper := widgets.NewDivWidget().SetClass("netmail-msg-view-header-wrapper").AddWidget(msgHeader)
	containerVBox.Add(msgHeaderWrapper)

	previewWidget := widgets.NewDivWidget().
		SetClass("netmail-msg-view-body").
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self NetmailViewAction) makeMessageHeaderRowSection(headerTable *widgets.TableWidget, name widgets.IWidget, value widgets.IWidget) {

	headerFromName := widgets.NewTableCellWidget()
	headerFromName.SetClass("netmail-msg-view-header-name")
	headerFromName.SetWidget(name)

	headerFromValue := widgets.NewTableCellWidget()
	headerFromValue.SetClass("netmail-msg-view-header-value")
	headerFromValue.SetWidget(value)

	headerTable.AddRow(
		widgets.NewTableRowWidget().
			AddCell(headerFromName).
			AddCell(headerFromValue),
	)

}

func (self NetmailViewAction) makeMessageHeaderSection(origMsg *mapper.NetmailMsg, msgBody *packet.MessageBody) widgets.IWidget {

	/* Make main header widget */
	headerTable := widgets.NewTableWidget().
		SetClass("netmail-msg-view-header")

	/* Make "From" section */
	var msgFrom string
	if origMsg.OrigAddr == "" {
		msgFrom = fmt.Sprintf("%s", origMsg.From)
	} else {
		msgFrom = fmt.Sprintf("%s (%s)", origMsg.From, origMsg.OrigAddr)
	}
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("From:"),
		widgets.NewTextWidgetWithText(msgFrom),
	)

	/* Make "To" section */
	var msgTo string
	if origMsg.DestAddr == "" {
		msgTo = fmt.Sprintf("%s", origMsg.To)
	} else {
		msgTo = fmt.Sprintf("%s (%s)", origMsg.To, origMsg.DestAddr)
	}
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("To:"),
		widgets.NewTextWidgetWithText(msgTo),
	)

	/* Make "Subject" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("Subject:"),
		widgets.NewTextWidgetWithText(origMsg.Subject),
	)

	/* Make "Date" section */
	newDate := fmt.Sprintf("%s", origMsg.DateWritten)
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("Date:"),
		widgets.NewTextWidgetWithText(newDate),
	)

	attachments := msgBody.GetAttachments()
	attachmentCount := len(attachments)
	if attachmentCount > 0 {

		attWidget := widgets.NewDivWidget()
		for idx, att := range attachments {

			attWidget.SetContent("📎")

			navigateAddr := fmt.Sprintf("/netmail/%s/attach/%d/view", origMsg.Hash, idx)

			navigateTitle := fmt.Sprintf("%s (%d kB)", att.GetName(), att.Len()/1024)
			navigateRow := widgets.NewLinkWidget().
				SetLink(navigateAddr).
				SetContent(navigateTitle)

			attWidget.AddWidget(navigateRow)

		}

		self.makeMessageHeaderRowSection(
			headerTable,
			widgets.NewTextWidgetWithText("Attachments:"),
			attWidget,
		)
	}

	return headerTable

}
