package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	commonfunc "github.com/vit1251/golden/internal/common"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/registry"
)

type NetmailAttachViewHandler struct {
	registry *registry.Container
}

func NewNetmailAttachViewHandler(registry *registry.Container) *NetmailAttachViewHandler {
	return &NetmailAttachViewHandler{
		registry: registry,
	}
}

func (self NetmailAttachViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	netmailMapper := mapperManager.GetNetmailMapper()

	/* Attach index */
	var attIdxParam string = r.PathValue("attidx")
	attIdx, _ := commonfunc.ParseSize([]byte(attIdxParam))

	//
	var msgid string = r.PathValue("msgid")
	origMsg, err3 := netmailMapper.GetMessageByHash(msgid)
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
	packedMessage := origMsg.GetPacket()

	bodyParser := packet.NewMessageBodyParser()
	bodyParser.SetDecodeAttachment(true)

	msgBody, _ := bodyParser.Parse(packedMessage)

	attachments := msgBody.GetAttachments()

	var attach *packet.MessageBodyAttach
	for aIndex, a := range attachments {
		if aIndex == attIdx {
			attach = packet.NewMessageBodyAttach()
			*attach = a
		}
	}

	if attach != nil {

		content := attach.GetData()
		attachReader := bytes.NewReader(content)
		io.Copy(w, attachReader)

	} else {
		response := fmt.Sprintf("No attach")
		http.Error(w, response, http.StatusNotFound)
		return
	}

}
