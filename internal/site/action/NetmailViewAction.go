package action

import (
	"fmt"
	"github.com/gorilla/mux"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/packet"
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

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
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
	bodyParser.SetDecodeAttachment(true)
	msgBody, _ := bodyParser.Parse(rawPacket)
	// TODO - use message parsing ... rawContent := msgBody.GetContent()
	// TODO - use message parsing ... content := string(rawContent)

	/* Processing message body */
	mtp := msg.NewMessageTextProcessor()
	doc, _ := mtp.Prepare(content)
	outDoc := doc.HTML()

	/* Update view counter */
	err5 := netmailMapper.ViewMessageByHash(msgHash)
	if err5 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash on netmailMapper: err = %+v", err5)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	/* Context actions */
	actionBar := self.renderActions(origMsg)
	containerVBox.Add(actionBar)

	msgHeader := self.makeMessageHeaderSection(origMsg, msgBody)
	msgHeaderWrapper := widgets2.NewDivWidget().SetClass("netmail-msg-view-header-wrapper").AddWidget(msgHeader)
	containerVBox.Add(msgHeaderWrapper)

	previewWidget := widgets2.NewDivWidget().
		SetClass("netmail-msg-view-body").
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self NetmailViewAction) makeMessageHeaderRowSection(headerTable *widgets2.TableWidget, name widgets2.IWidget, value widgets2.IWidget) {

	headerFromName := widgets2.NewTableCellWidget()
	headerFromName.SetClass("netmail-msg-view-header-name")
	headerFromName.SetWidget(name)

	headerFromValue := widgets2.NewTableCellWidget()
	headerFromValue.SetClass("netmail-msg-view-header-value")
	headerFromValue.SetWidget(value)

	headerTable.AddRow(
		widgets2.NewTableRowWidget().
			AddCell(headerFromName).
			AddCell(headerFromValue),
	)

}

func (self NetmailViewAction) makeMessageHeaderSection(origMsg *mapper.NetmailMsg, msgBody *packet.MessageBody) widgets2.IWidget {

	/* Make main header widget */
	headerTable := widgets2.NewTableWidget().
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
		widgets2.NewTextWidgetWithText("From:"),
		widgets2.NewTextWidgetWithText(msgFrom),
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
		widgets2.NewTextWidgetWithText("To:"),
		widgets2.NewTextWidgetWithText(msgTo),
	)

	/* Make "Subject" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets2.NewTextWidgetWithText("Subject:"),
		widgets2.NewTextWidgetWithText(origMsg.Subject),
	)

	/* Make "Date" section */
	newDate := fmt.Sprintf("%s", origMsg.DateWritten)
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets2.NewTextWidgetWithText("Date:"),
		widgets2.NewTextWidgetWithText(newDate),
	)

	attachments := msgBody.GetAttachments()
	attachmentCount := len(attachments)
	if attachmentCount > 0 {

		attWidget := widgets2.NewDivWidget()
		for idx, att := range attachments {

			attWidget.SetContent("📎")

			navigateAddr := fmt.Sprintf("/netmail/%s/attach/%d/view", origMsg.Hash, idx)

			navigateTitle := fmt.Sprintf("%s (%d kB)", att.GetName(), att.Len()/1024)
			navigateRow := widgets2.NewLinkWidget().
				SetLink(navigateAddr).
				SetContent(navigateTitle)

			attWidget.AddWidget(navigateRow)

		}

		self.makeMessageHeaderRowSection(
			headerTable,
			widgets2.NewTextWidgetWithText("Attachments:"),
			attWidget,
		)
	}

	return headerTable

}

func (self NetmailViewAction) renderActions(origMsg *mapper.NetmailMsg) widgets2.IWidget {
	actionBar := widgets2.NewActionMenuWidget()

	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/netmail/%s/reply", origMsg.Hash)).
		SetClass("mr-2").
		SetIcon("icofont-edit").
		SetLabel("Reply"))

	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/netmail/%s/dump", origMsg.Hash)).
		SetClass("mr-2").
		SetIcon("icofont-dump").
		SetLabel("Info"))

	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/netmail/%s/remove", origMsg.Hash)).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Delete"))

	return actionBar
}
