package action

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type NetmailRemoveApiAction struct {
	Action
}

func NewNetmailRemoveApiAction() *NetmailRemoveApiAction {
	return new(NetmailRemoveApiAction)
}

func (self *NetmailRemoveApiAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	netmailMapper := mapperManager.GetNetmailMapper()

	var code int = 0

	r.ParseForm()
	msgHash := r.PostForm.Get("msgHash")

	err := netmailMapper.RemoveMessageByHash(msgHash)
	if err != nil {
		code = 1
	}

	p := make(map[string]interface{})
	p["code"] = code
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
