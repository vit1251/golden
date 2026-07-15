package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoAreaIndexHandler struct {
    registry *registry.Container
}

func NewEchoAreaIndexHandler(registry *registry.Container) *EchoAreaIndexHandler {
    return &EchoAreaIndexHandler{
	registry: registry,
    }
}

func (h *EchoAreaIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()

    areas, err := echoAreaMapper.GetAreas()
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }

    var areaHeaders []views.AreaHeader
    for _, a := range areas {
	areaHeaders = append(areaHeaders, views.AreaHeader{
	    Name:        a.GetName(),
	    Summary:     a.GetSummary(),
	    IndexURL:    "/echo/" + a.GetAreaIndex(),
	    NewMsgCount: a.GetNewMessageCount(),
	})
    }

    data := views.EchoAreaIndexData{
	Actions: []views.ToolbarAction{
	    {Label: "Create", URL: "/echo/create", Icon: "edit"},
	},
	Areas: areaHeaders,
    }

    err = views.Page("Echo Areas", views.EchoAreaIndexView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

