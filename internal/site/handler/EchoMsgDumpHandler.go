package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
    "github.com/vit1251/golden/internal/utils"
)

type EchoMsgDumpHandler struct {
    registry *registry.Container
}

func NewEchoMsgDumpHandler(registry *registry.Container) *EchoMsgDumpHandler {
    return &EchoMsgDumpHandler{
	registry: registry,
    }
}

func (h *EchoMsgDumpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()
    echoMapper := mapperManager.GetEchoMapper()

    areaIndex := r.PathValue("echoname")
    area, err := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    msgHash := r.PathValue("msgid")
    origMsg, err := echoMapper.GetMessageByHash(area.GetName(), msgHash)
    if err != nil {
        http.Error(w, "Message not found", http.StatusInternalServerError)
        return
    }

    dump := utils.HexDumpGrouped(origMsg.Packet)

    data := views.EchoMsgDumpData{
        Actions: []views.ToolbarAction{
            {Label: "Back",   URL: "/echo/" + areaIndex, Icon: "arrow-left"},
        },
        Dump: dump,
    }

    err = views.Page(area.GetName(), views.EchoMsgDumpView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
