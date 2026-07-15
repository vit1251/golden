package handler

import (
	"net/http"

	"github.com/vit1251/golden/internal/site/views"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoAreaUpdateHandler struct {
    registry *registry.Container
}

func NewEchoAreaUpdateHandler(registry *registry.Container) *EchoAreaUpdateHandler {
    return &EchoAreaUpdateHandler{
	registry: registry,
    }
}

func (u *EchoAreaUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(u.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()

    //
    var areaIndex string = r.PathValue("echoname")
    area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err1 != nil {
	http.Error(w, err1.Error(), http.StatusInternalServerError)
	return
    }

    data := views.EchoAreaUpdateData{
        Actions: []views.ToolbarAction{
            {Label: "Back", URL: "/echo/" + area.GetAreaIndex(), Icon: "arrow-left"},
            {Label: "Purge", URL: "/echo/" + area.GetAreaIndex() + "/purge", Icon: "trash-2"},
        },
        ActionURL: "/echo/" + area.GetAreaIndex() + "/update",
        AreaName:  area.GetName(),
        Summary:   area.Summary,
        Charset:   area.GetCharset(),
        SortOrder: area.GetOrder(),
    }
    err := views.Page("Edit area", views.EchoAreaUpdateView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
