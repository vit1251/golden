package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/echomail"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/widgets"
	"html/template"
	"net/http"
)

type EchoMsgViewAction struct {
	Action
}

func NewEchoMsgViewAction() *EchoMsgViewAction {
	va := new(EchoMsgViewAction)
	return va
}

func (self *EchoMsgViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Restore managers */
	areaManager := self.restoreAreaManager()
	messageManager := self.restoreMessageManager()

	/* Get "echoname" in user request */
	vars := mux.Vars(r)

	/* Restore area by "echoname" key */
	echoTag := vars["echoname"]
	//log.Printf("echoTag = %+v", echoTag)
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName in AreaManager: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Restore message by "echoname" and "msgid" key */
	msgHash := vars["msgid"]
	//log.Printf("msgid = %+v", msgid)
	origMsg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash in MessageManager: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if origMsg == nil {
		response := fmt.Sprintf("Fail on GetMessageByHash in MessageManager: err = %+v", fmt.Errorf("origMsg is emprty"))
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Get message body */
	content := origMsg.GetContent()

	/* Prepare HTML message body */
	mtp := msg.NewMessageTextProcessor()
	if err4 := mtp.Prepare(content); err4 != nil {
		response := fmt.Sprintf("Fail on Prepare in MessageTextProcessor: err = %+v", err4)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	outDoc := mtp.HTML()

	/* Update message view counter */
	err5 := messageManager.ViewMessageByHash(echoTag, msgHash)
	if err5 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash on messageManager: err = %+v", err5)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Make HTML page content */
	mainWidget := self.makeMainEchoMsgViewWidget(area, origMsg, outDoc)

	/* Render process */
	if err := mainWidget.Render(w); err != nil {
		status := fmt.Sprintf("Fail on Render in Widget: err = %+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}


func (self EchoMsgViewAction) makeMessageHeaderRowSection(headerTable *widgets.TableWidget, name widgets.IWidget, value widgets.IWidget) {

	headerFromName := widgets.NewTableCellWidget()
	headerFromName.SetClass("echo-msg-view-header-name")
	headerFromName.SetWidget(name)

	headerFromValue := widgets.NewTableCellWidget()
	headerFromValue.SetClass("echo-msg-view-header-value")
	headerFromValue.SetWidget(value)

	headerTable.AddRow(
		widgets.NewTableRowWidget().
			AddCell(headerFromName).
			AddCell(headerFromValue),
	)

}

func (self EchoMsgViewAction) makeMessageHeaderSection(origMsg msg.Message) widgets.IWidget {

	/* Make main header widget */
	headerTable := widgets.NewTableWidget().
		SetClass("echo-msg-view-header")

	/* Make "Area" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("Area:"),
		widgets.NewTextWidgetWithText(origMsg.Area),
	)

	/* Make "From" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("From:"),
		widgets.NewTextWidgetWithText(origMsg.From),
	)

	/* Make "To" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets.NewTextWidgetWithText("To:"),
		widgets.NewTextWidgetWithText(origMsg.To),
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

	return headerTable

}

func (self EchoMsgViewAction) makeMainEchoMsgViewWidget(area *echomail.Area, origMsg *msg.Message, outDoc template.HTML) widgets.IWidget {

	mainWidget := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	mainWidget.SetWidget(vBox)

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
			SetLink(fmt.Sprintf("/echo/%s//message/%s/reply", area.Name(), origMsg.Hash)).
			SetIcon("icofont-edit").
			SetLabel("Reply")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/message/%s/remove", area.Name(), origMsg.Hash)).
			SetIcon("icofont-remove").
			SetLabel("Delete"))
	containerVBox.Add(amw)

	/* Message header section */
	msgHeader := self.makeMessageHeaderSection(*origMsg)
	msgHeaderWrapper := widgets.NewDivWidget().SetClass("echo-msg-view-header-wrapper").SetWidget(msgHeader)
	containerVBox.Add(msgHeaderWrapper)

	/* Message body */
	previewWidget := widgets.NewDivWidget().
		SetClass("echo-msg-view-body").
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	return mainWidget

}
