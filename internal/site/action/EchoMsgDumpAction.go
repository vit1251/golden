package action

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/mux"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
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

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
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
	actionsBar := self.renderActions(area, origMsg)
	containerVBox.Add(actionsBar)

	indexTable := widgets2.NewTableWidget().
		SetClass("table")

	containerVBox.Add(indexTable)

	previewWidget := widgets2.NewPreWidget().
		SetContent(string(outDoc))
	containerVBox.Add(previewWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *EchoMsgDumpAction) renderActions(area *mapper.Area, origMsg *msg.Message) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())

	actionBar := widgets2.NewActionMenuWidget()

	/* Reply */
	replyAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/reply").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(replyAddr).
		SetIcon("icofont-edit").
		SetLabel("Reply"))

	/* Remove */
	removeAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/remove").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", origMsg.Hash).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(removeAddr).
		SetIcon("icofont-remove").
		SetLabel("Delete"))

	return actionBar
}
