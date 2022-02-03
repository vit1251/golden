package action

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/mux"
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
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := echoAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	msgHeaders, err112 := echoMapper.GetMessageHeaders(echoTag)
	if err112 != nil {
		response := fmt.Sprintf("Fail on GetAreas")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeaders = %+v", msgHeaders)

	//
	msgHash := vars["msgid"]
	origMsg, err3 := echoMapper.GetMessageByHash(echoTag, msgHash)
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
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s//message/%s/reply", area.GetName(), origMsg.Hash)).
			SetIcon("icofont-edit").
			SetLabel("Reply")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/message/%s/remove", area.GetName(), origMsg.Hash)).
			SetIcon("icofont-remove").
			SetLabel("Delete"))
	containerVBox.Add(amw)

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
