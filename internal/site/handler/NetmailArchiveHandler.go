package handler

import (
    "log"
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type NetmailArchiveHandler struct {
    registry *registry.Container
}

func NewNetmailArchiveHandler(registry *registry.Container) *NetmailArchiveHandler {
    return &NetmailArchiveHandler{
	registry: registry,
    }
}

func (h *NetmailArchiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    netmailMapper := mapperManager.GetNetmailMapper()

    //
    msgHash := r.PathValue("msgid")
    msg, err3 := netmailMapper.GetMessageByHash(msgHash)
    if err3 != nil {
	http.Error(w, "Fail on GetMessageByHash", http.StatusInternalServerError)
	return
    }
    log.Printf("msg = %+v", msg)

    if err1 := netmailMapper.ArchiveMessageByHash(msgHash); err1 != nil {
	http.Error(w, "Fail on GetMessageByHash", http.StatusInternalServerError)
	return
    }

    /* Redirect */
    http.Redirect(w, r, "/netmail", 303)

}
