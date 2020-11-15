package action

import (
	"fmt"
	"github.com/gorilla/mux"
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


	fmt.Printf("packedMessage = %+v", packedMessage)

}
