package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"net/http"
)

type APIStatAction struct {
	APIAction
}

func NewAPIStatAction() *APIStatAction {
	smac := new(APIStatAction)
	return smac
}

func (self *APIStatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Calculate summary */
	var newDirectMsgCount int
	var newEchoMsgCount int
	var newFileCount int
	self.Container.Invoke(func(nm *netmail.NetmailManager, em *msg.MessageManager, fm *file.FileManager) {
		newDirectMsgCount, _ = nm.GetMessageNewCount()
		newEchoMsgCount, _ = em.GetMessageNewCount()
		newFileCount, _ = fm.GetMessageNewCount()
	})

	p := make(map[string]interface{})
	p["EchomailMessageCount"] = newEchoMsgCount
	p["NetmailMessageCount"] = newDirectMsgCount
	p["FileCount"] = newFileCount
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
