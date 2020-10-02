package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"net/http"
)

type StatAction struct {
	Action
}

func NewStatAction() *StatAction {
	smac := new(StatAction)
	return smac
}

func (self *StatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
	p["code"] = 0
	p["EchomailMessageCount"] = newEchoMsgCount
	p["NetmailMessageCount"] = newDirectMsgCount
	p["FileCount"] = newFileCount
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
