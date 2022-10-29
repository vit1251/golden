package action

import (
	"fmt"
	"github.com/gorilla/mux"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"html/template"
	"log"
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

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Parse URL parameters */
	vars := mux.Vars(r)

	//
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName in echoAreaMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Restore message by "echoname" and "msgid" key */
	msgHash := vars["msgid"]
	log.Printf("msgHash = %+v", msgHash)

	var areaName string = area.GetName()
	origMsg, err3 := echoMapper.GetMessageByHash(areaName, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash in echoMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if origMsg == nil {
		response := fmt.Sprintf("Fail on GetMessageByHash in echoMapper: err = %+v", fmt.Errorf("origMsg is emprty"))
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Get message body */
	content := origMsg.GetContent()

	/* Prepare HTML message body */
	mtp := msg.NewMessageTextProcessor()
	doc, _ := mtp.Prepare(content)
	outDoc := doc.HTML()

	/* Update message view counter */
	err5 := echoMapper.ViewMessageByHash(areaName, msgHash)
	if err5 != nil {
		response := fmt.Sprintf("Fail on ViewMessageByHash on echoMapper: err = %+v", err5)
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

func (self EchoMsgViewAction) makeMessageHeaderRowSection(headerTable *widgets2.TableWidget, name widgets2.IWidget, value widgets2.IWidget) {

	headerFromName := widgets2.NewTableCellWidget()
	headerFromName.SetClass("echo-msg-view-header-name")
	headerFromName.SetWidget(name)

	headerFromValue := widgets2.NewTableCellWidget()
	headerFromValue.SetClass("echo-msg-view-header-value")
	headerFromValue.SetWidget(value)

	headerTable.AddRow(
		widgets2.NewTableRowWidget().
			AddCell(headerFromName).
			AddCell(headerFromValue),
	)

}

func (self EchoMsgViewAction) makeMessageHeaderSection(origMsg msg.Message) widgets2.IWidget {

	/* Make main header widget */
	headerTable := widgets2.NewTableWidget().
		SetClass("echo-msg-view-header")

	/* Make "Area" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets2.NewTextWidgetWithText("Area:"),
		widgets2.NewTextWidgetWithText(origMsg.Area),
	)

	/* Make "From" section */
	var newFrom string
	if origMsg.FromAddr != "" {
		newFrom = fmt.Sprintf("%s ( %s )", origMsg.From, origMsg.FromAddr)
	} else {
		newFrom = origMsg.From
	}
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets2.NewTextWidgetWithText("From:"),
		widgets2.NewTextWidgetWithText(newFrom),
	)

	/* Make "To" section */
	self.makeMessageHeaderRowSection(
		headerTable,
		widgets2.NewTextWidgetWithText("To:"),
		widgets2.NewTextWidgetWithText(origMsg.To),
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

	return headerTable

}

func (self EchoMsgViewAction) makeMainEchoMsgViewWidget(area *mapper.Area, origMsg *msg.Message, outDoc template.HTML) widgets2.IWidget {

	mainWidget := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	mainWidget.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")
	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()
	container.AddWidget(containerVBox)

	/* Context actions */
	actionsBar := self.renderActions(area, origMsg)
	containerVBox.Add(actionsBar)

	/* Message header section */
	msgHeader := self.makeMessageHeaderSection(*origMsg)
	msgHeaderWrapper := widgets2.NewDivWidget().SetClass("echo-msg-view-header-wrapper").AddWidget(msgHeader)
	containerVBox.Add(msgHeaderWrapper)

	/* Message body */
	previewWidget := widgets2.NewDivWidget().
		SetClass("echo-msg-view-body").
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	return mainWidget

}

func (self *EchoMsgViewAction) renderActions(area *mapper.Area, origMsg *msg.Message) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	actionBar := widgets2.NewActionMenuWidget()

	/* Reply */
	messageReplyAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/reply").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(messageReplyAddr).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel("Reply"))

	/* Action Remove */
	messageRemoveAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/remove").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(messageRemoveAddr).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Remove"))

	/* Action Dump */
	messageDumpAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/dump").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(messageDumpAddr).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Dump"))

	/* Action Twit */
	messageTwitAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/twit").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(messageTwitAddr).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Twit"))

	return actionBar
}
