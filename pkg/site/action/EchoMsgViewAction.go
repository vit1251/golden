package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/widgets"
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

	areaManager := self.restoreAreaManager()
	messageManager := self.restoreMessageManager()
	//configManager := self.restoreConfigManager()

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	msgHeaders, err112 := messageManager.GetMessageHeaders(echoTag)
	if err112 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeaders = %+v", msgHeaders)

	//
	msgHash := vars["msgid"]
	origMsg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
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

	fmt.Printf("orgMsg = %+v", origMsg)

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
	err5 := messageManager.ViewMessageByHash(echoTag, msgHash)
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

	/* Render process */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
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
		widgets.NewTextWidgetWithText(origMsg.To),
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
