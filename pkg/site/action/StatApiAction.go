package action

import (
	"encoding/json"
	"net/http"
)

type StatApiAction struct {
	Action
}

func NewStatApiAction() *StatApiAction {
	smac := new(StatApiAction)
	return smac
}

func (self *StatApiAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	netmailManager := self.restoreNetmailManager()
	messageManager := self.restoreMessageManager()
	fileManager := self.restoreFileManager()

	/* Calculate summary */
	var newDirectMsgCount int
	var newEchoMsgCount int
	var newFileCount int

	newEchoMsgCount, _ = messageManager.GetMessageNewCount()
	newDirectMsgCount, _ = netmailManager.GetMessageNewCount()
	newFileCount, _ = fileManager.GetMessageNewCount()

	p := make(map[string]interface{})
	p["code"] = 0
	p["EchomailMessageCount"] = newEchoMsgCount
	p["NetmailMessageCount"] = newDirectMsgCount
	p["FileCount"] = newFileCount
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
