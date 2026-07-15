package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type NetmailIndexHandler struct {
    registry *registry.Container
}

func NewNetmailIndexHandler(registry *registry.Container) *NetmailIndexHandler {
    return &NetmailIndexHandler{
	registry: registry,
    }
}

func (self NetmailIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // Шаг 0. Получаем необходимые источники данных
    mapperManager := mapper.RestoreMapperManager(self.registry)
    netmailMapper := mapperManager.GetNetmailMapper()

    // Шаг 1. Получаем данные
    msgHeaders, err1 := netmailMapper.GetMessageHeaders()
    if err1 != nil {
	http.Error(w, "Fail on GetMessageHeaders", http.StatusInternalServerError)
	return
    }

    var msgs []views.NetmailMsgHeader
    for _, m := range msgHeaders {
	msgs = append(msgs, views.NetmailMsgHeader{
	    Hash:    m.Hash,
	    From:    m.From,
	    Subject: m.Subject,
	    Date:    m.DateWritten.Format("2006-01-02 15:04"),
	    IsNew:   m.ViewCount == 0,
	    ViewURL: "/netmail/" + m.Hash + "/view",
	})
    }
    data := views.NetmailIndexData{
	Actions: []views.ToolbarAction{
	    {Label: "New", URL: "/netmail/compose", Icon: "edit"},
	},
	Messages: msgs,
    }
    err := views.Page("Netmail", views.NetmailIndexView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
