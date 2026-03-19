package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type NetmailRemoveApiHandler struct {
	registry *registry.Container
}

func NewNetmailRemoveApiHandler(registry *registry.Container) *NetmailRemoveApiHandler {
	return &NetmailRemoveApiHandler{
		registry: registry,
	}
}

func (self *NetmailRemoveApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
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
