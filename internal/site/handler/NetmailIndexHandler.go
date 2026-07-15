package handler

import (
    "strconv"
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
    page := 1
    if p := r.URL.Query().Get("page"); p != "" {
	page, _ = strconv.Atoi(p)
	if page < 1 { page = 1 }
    }
    const limit = 20
    offset := (page - 1) * limit

    msgHeaders, err1 := netmailMapper.GetMessageHeadersPage(limit, offset)
    if err1 != nil {
	http.Error(w, "Fail on GetMessageHeaders", http.StatusInternalServerError)
	return
    }
    totalCount, err2 := netmailMapper.GetMessageCount()
    if err2 != nil {
	http.Error(w, "Fail on GetMessageCount", http.StatusInternalServerError)
	return
    }
    totalPages := (totalCount + limit - 1) / limit

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
	    {Label: "Compose", URL: "/netmail/compose", Icon: "edit"},
	},
	Messages: msgs,
	Pagination: views.PaginationData{
    	    CurrentPage: page,
    	    TotalPages:  totalPages,
	    BaseURL:     "/netmail",
        },
    }
    err := views.Page("Netmail", views.NetmailIndexView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
