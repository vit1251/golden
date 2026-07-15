package handler

import (
    "net/http"
    "strconv"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoMsgIndexHandler struct {
    registry *registry.Container
}

func NewEchoMsgIndexHandler(registry *registry.Container) *EchoMsgIndexHandler {
    return &EchoMsgIndexHandler{
	registry: registry,
    }
}

func (h *EchoMsgIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()
    echoMapper := mapperManager.GetEchoMapper()

    areaIndex := r.PathValue("echoname")

    area, err := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }
    areaName := area.GetName()

    // Pagination
    page := 1
    if p := r.URL.Query().Get("page"); p != "" {
	page, _ = strconv.Atoi(p)
	if page < 1 {
	    page = 1
	}
    }

    const limit = 20
    offset := (page - 1) * limit

    headers, err := echoMapper.GetMessageHeadersPage(areaName, limit, offset)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }

    totalCount, err := echoMapper.GetMessageCount(areaName)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }
    totalPages := (totalCount + limit - 1) / limit

    // Маппим в view-структуры
    var msgs []views.MsgHeader
    for _, m := range headers {
	msgs = append(msgs, views.MsgHeader{
	    Hash:    m.Hash,
	    From:    m.From,
	    To:      m.To,
	    Subject: m.Subject,
	    Date:    m.DateWritten.Format("2006-01-02 15:04"),
	    IsNew:   m.ViewCount == 0,
	    ViewURL: "/echo/" + areaIndex + "/message/" + m.Hash + "/view",
	})
    }

    data := views.EchoMsgIndexData{
        Actions: []views.ToolbarAction{
            {Label: "Back", URL: "/echo", Icon: "arrow-left"},
            {Label: "Compose", URL: "/echo/" + areaIndex + "/message/compose", Icon: "edit"},
	    {Label: "Tree", URL: "/echo/" + areaIndex + "/tree", Icon: "tree"},
	    {Label: "Mark read", URL: "/echo/" + areaIndex + "/mark", Icon: "check"},
	    {Label: "Edit", URL: "/echo/" + areaIndex + "/update", Icon: "settings"},
        },
	AreaName: areaName,
	Messages: msgs,
	Pagination: views.PaginationData{
	    CurrentPage: page,
	    TotalPages:  totalPages,
	    BaseURL:     "/echo/" + areaIndex,
	},
    }

    err = views.Page(areaName, views.EchoMsgIndexView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
