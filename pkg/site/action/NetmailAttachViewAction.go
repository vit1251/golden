package action

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	commonfunc "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/packet"
	"io"
	"net/http"
)

type NetmailAttachViewAction struct {
	Action
}

func NewNetmailAttachViewAction() *NetmailAttachViewAction {
	va := new(NetmailAttachViewAction)
	return va
}

func (self NetmailAttachViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	netmailMapper := mapperManager.GetNetmailMapper()

	//
	vars := mux.Vars(r)

	/* Attach index */
	attIdxParam := vars["attidx"]
	attIdx, _ := commonfunc.ParseSize([]byte(attIdxParam))

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
