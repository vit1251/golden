package action

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type NetmailDumpAction struct {
	Action
}

func NewNetmailDumpAction() *NetmailDumpAction {
	va := new(NetmailDumpAction)
	return va
}

func (self NetmailDumpAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
			SetLink(fmt.Sprintf("/netmail/%s/view", origMsg.Hash)).
			SetClass("netmail-reply-action").
			SetIcon("icofont-view").
			SetLabel("View"))

	containerVBox.Add(amw)

	/* Dump message */
	rawPacket := origMsg.GetPacket()
	outDoc := hex.Dump(rawPacket)

	//
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
