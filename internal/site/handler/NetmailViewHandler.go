package handler

import (
    "fmt"
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/msg"
    "github.com/vit1251/golden/pkg/packet"
    "github.com/vit1251/golden/pkg/registry"
)

type NetmailViewHandler struct {
	registry *registry.Container
}

func NewNetmailViewHandler(registry *registry.Container) *NetmailViewHandler {
	return &NetmailViewHandler{
		registry: registry,
	}
}

func (h *NetmailViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
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
    content := origMsg.GetContent()

    /* Preprocess message body (attachments) */
    rawPacket := origMsg.GetPacket()
    bodyParser := packet.NewMessageBodyParser()
    bodyParser.SetDecodeAttachment(true)
    msgBody, _ := bodyParser.Parse(rawPacket)
    // TODO - use message parsing ... rawContent := msgBody.GetContent()
    // TODO - use message parsing ... content := string(rawContent)

    /* Processing message body */
    mtp := msg.NewMessageTextProcessor()
    doc := mtp.Prepare(content)
    outDoc := doc.HTML()

    /* Update view counter */
    err5 := netmailMapper.ViewMessageByHash(msgHash)
    if err5 != nil {
	response := fmt.Sprintf("Fail on ViewMessageByHash on netmailMapper: err = %+v", err5)
	http.Error(w, response, http.StatusInternalServerError)
	return
    }

    var msgFrom string
    if origMsg.OrigAddr == "" {
	msgFrom = origMsg.From
    } else {
	msgFrom = fmt.Sprintf("%s (%s)", origMsg.From, origMsg.OrigAddr)
    }
    var msgTo string
    if origMsg.DestAddr == "" {
	msgTo = origMsg.To
    } else {
        msgTo = fmt.Sprintf("%s (%s)", origMsg.To, origMsg.DestAddr)
    }

    var atts []views.NetmailAttachment
    attachments := msgBody.GetAttachments()
    for idx, att := range attachments {
	atts = append(atts, views.NetmailAttachment{
    	    Name: att.GetName(),
    	    URL:  fmt.Sprintf("/netmail/%s/attach/%d/view", origMsg.Hash, idx),
    	    Size: fmt.Sprintf("%d kB", att.Len()/1024),
	})
    }

    // Шаг 2. Рендеринг
    data := views.NetmailViewData{
        Actions: []views.ToolbarAction{
            {Label: "Back",    URL: "/netmail",                                  Icon: "arrow-left"},
            {Label: "Reply",   URL: "/netmail/" + origMsg.Hash + "/reply",       Icon: "edit"},
            {Label: "Dump",    URL: "/netmail/" + origMsg.Hash + "/dump",        Icon: "file-code"},
            {Label: "Archive", URL: "/netmail/" + origMsg.Hash + "/archive",     Icon: "archive"},
        },
        From:    msgFrom,
        To:      msgTo,
        Subject: origMsg.Subject,
        Date:    origMsg.DateWritten.Format("2006-01-02 15:04"),
        Body:    string(outDoc),
        Attachments: atts,
    }
    err := views.Page("Netmail", views.NetmailViewView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
