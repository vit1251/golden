package handler

import (
    "fmt"
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/msg"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoMsgViewHandler struct {
    registry *registry.Container
}

func NewEchoMsgViewHandler(registry *registry.Container) *EchoMsgViewHandler {
    return &EchoMsgViewHandler{
	registry: registry,
    }
}

func (h *EchoMsgViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()
    echoMapper := mapperManager.GetEchoMapper()

    areaIndex := r.PathValue("echoname")
    area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    msgid := r.PathValue("msgid")
    origMsg, err2 := echoMapper.GetMessageByHash(area.GetName(), msgid)
    if err2 != nil || origMsg == nil {
        http.Error(w, "Message not found", http.StatusInternalServerError)
        return
    }

    content := origMsg.GetContent()
    mtp := msg.NewMessageTextProcessor()
    doc := mtp.Prepare(content)

    echoMapper.ViewMessageByHash(area.GetName(), msgid)

    var newFrom string
    if origMsg.FromAddr != "" {
        newFrom = fmt.Sprintf("%s ( %s )", origMsg.From, origMsg.FromAddr)
    } else {
        newFrom = origMsg.From
    }

    data := views.EchoMsgViewData{
        Actions: []views.ToolbarAction{
            {Label: "Back",   URL: "/echo/" + areaIndex, Icon: "arrow-left"},
            {Label: "Reply",  URL: "/echo/" + areaIndex + "/message/" + msgid + "/reply", Icon: "edit"},
            {Label: "Archive", URL: "/echo/" + areaIndex + "/message/" + msgid + "/archive", Icon: "archive"},
            {Label: "Dump",   URL: "/echo/" + areaIndex + "/message/" + msgid + "/dump", Icon: "file-code"},
        },
        AreaName: area.GetName(),
        From:     newFrom,
        To:       origMsg.To,
        Subject:  origMsg.Subject,
        Date:     origMsg.DateWritten.Format("2006-01-02 15:04"),
        Body:     string(doc.HTML()),
    }

    err := views.Page(area.GetName(), views.EchoMsgViewView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
