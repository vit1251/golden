package action

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type EchoMsgDumpAction struct {
	Action
}

func NewEchoMsgDumpAction() *EchoMsgDumpAction {
	va := new(EchoMsgDumpAction)
	return va
}

func (self *EchoMsgDumpAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	//
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	var areaName string = area.GetName()
	msgHeaders, err112 := echoMapper.GetMessageHeaders(areaName)
	if err112 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeaders = %+v", msgHeaders)

	//
	msgHash := vars["msgid"]
	origMsg, err3 := echoMapper.GetMessageByHash(areaName, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	//outDoc := origMsg.GetContent()
	outDoc := hex.Dump(origMsg.Packet)

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
	actionsBar := self.renderActions(area, origMsg)
	containerVBox.Add(actionsBar)

	indexTable := widgets.NewTableWidget().
		SetClass("table")

	containerVBox.Add(indexTable)

	previewWidget := widgets.NewPreWidget().
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *EchoMsgDumpAction) renderActions(area *mapper.Area, origMsg *msg.Message) widgets.IWidget {

	urlManager := self.restoreUrlManager()

	actionBar := widgets.NewActionMenuWidget()

	/* Reply */
	replyAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/reply").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets.NewMenuAction().
		SetLink(replyAddr).
		SetIcon("icofont-edit").
		SetLabel("Reply"))

	/* Remove */
	removeAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/remove").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets.NewMenuAction().
		SetLink(removeAddr).
		SetIcon("icofont-remove").
		SetLabel("Delete"))

	return actionBar
}
