package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/audio"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoViewAction struct {
	Action
}

func NewEchoViewAction() *EchoViewAction {
	va := new(EchoViewAction)
	return va
}

func (self *EchoViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *msg.AreaManager
	var messageManager *msg.MessageManager
	var configManager *setup.ConfigManager
	self.Container.Invoke(func(am *msg.AreaManager, mm *msg.MessageManager, cm *setup.ConfigManager) {
		areaManager = am
		messageManager = mm
		configManager = cm
	})

	//
	realName, _ := configManager.Get("main", "RealName", "")

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
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Play */
	if origMsg.To == realName {
		am := audio.NewAudioManager()
		am.Play("short-notice.mp3")
	}

	var content string
	if origMsg != nil {
		content = origMsg.GetContent()
	} else {
		content = "!! Unable restore message !!"
	}
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

	mmw := widgets.NewMainMenuWidget()
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

	indexTable := widgets.NewTableWidget().
		SetClass("table")

	//                <div>{{ .Msg.From }}</div>
	//                <div>{{ .Msg.To }}</div>
	//                <div>{{ .Msg.Subject }}</div>


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
