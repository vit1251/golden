package handler

import (
	"encoding/hex"
	"fmt"
	"net/http"

	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type NetmailDumpHandler struct {
	registry *registry.Container
}

func NewNetmailDumpHandler(registry *registry.Container) *NetmailDumpHandler {
	return &NetmailDumpHandler{
		registry: registry,
	}
}

func (self NetmailDumpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	netmailMapper := mapperManager.GetNetmailMapper()

	//
	var msgHash string = r.PathValue("msgid")
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

	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets2.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	/* Context handlers */
	actionBar := self.renderHandlers(origMsg)
	containerVBox.Add(actionBar)

	/* Dump message */
	rawPacket := origMsg.GetPacket()
	outDoc := hex.Dump(rawPacket)

	//
	previewWidget := widgets2.NewPreWidget().
		SetContent(string(outDoc))

	containerVBox.Add(previewWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self NetmailDumpHandler) renderHandlers(origMsg *mapper.NetmailMsg) widgets2.IWidget {

	actionBar := widgets2.NewActionMenuWidget()

	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/netmail/%s/view", origMsg.Hash)).
		SetIcon("icofont-view").
		SetClass("mr-2").
		SetLabel("View"))

	return actionBar
}
