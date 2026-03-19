package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type NetmailRemoveHandler struct {
	registry *registry.Container
}

func NewNetmailRemoveHandler(registry *registry.Container) *NetmailRemoveHandler {
	return &NetmailRemoveHandler{
		registry: registry,
	}
}

func (self *NetmailRemoveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	netmailMapper := mapperManager.GetNetmailMapper()

	//
	var msgHash string = r.PathValue("msgid")
	_, err3 := netmailMapper.GetMessageByHash(msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	log.Printf("msg = %s\n", msgHash)

	if err := netmailMapper.RemoveMessageByHash(msgHash); err != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/netmail")
	http.Redirect(w, r, newLocation, 303)

}
