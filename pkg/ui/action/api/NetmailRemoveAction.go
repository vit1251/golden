package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/netmail"
	"net/http"
)

type NetmailRemoveAction struct {
	Action
}

func NewNetmailRemoveAction() *NetmailRemoveAction {
	return new(NetmailRemoveAction)
}

func (self *NetmailRemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var code int

	var netmailManager *netmail.NetmailManager
	self.Container.Invoke(func(nm *netmail.NetmailManager) {
		netmailManager = nm
	})

	r.ParseForm()
	msgHash := r.PostForm.Get("msgHash")

	err := netmailManager.RemoveMessageByHash(msgHash)
	if err != nil {
		code = 1
	}

	p := make(map[string]interface{})
	p["code"] = code
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
