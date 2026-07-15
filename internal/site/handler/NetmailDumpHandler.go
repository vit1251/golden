package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/utils"
    "github.com/vit1251/golden/internal/site/views"
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

func (h *NetmailDumpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    netmailMapper := mapperManager.GetNetmailMapper()

    msgHash := r.PathValue("msgid")
    origMsg, err := netmailMapper.GetMessageByHash(msgHash)
    if err != nil || origMsg == nil {
        http.Error(w, "Message not found", http.StatusInternalServerError)
        return
    }

    dump := utils.HexDumpGrouped(origMsg.GetPacket())

    data := views.EchoMsgDumpData{
        Actions: []views.ToolbarAction{
            {Label: "Back",   URL: "/netmail/" + msgHash + "/view", Icon: "arrow-left"},
        },
        Dump: dump,
    }

    err = views.Page("Netmail", views.EchoMsgDumpView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

