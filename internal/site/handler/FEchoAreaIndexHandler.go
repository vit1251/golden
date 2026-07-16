package handler

import (
    "fmt"
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoAreaIndexHandler struct {
    registry *registry.Container
}

func NewFEchoAreaIndexHandler(registry *registry.Container) *FEchoAreaIndexHandler {
    return &FEchoAreaIndexHandler{
	registry: registry,
    }
}

func (h *FEchoAreaIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // Шаг 0. Подготовка служб
    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()

    // Шаг 1. Получение данных из хранилища
    areas, err1 := fileAreaMapper.GetAreas()
    if err1 != nil {
	response := fmt.Sprintf("Fail in GetAreas on fileAreaMapper: err = %+v", err1)
	http.Error(w, response, http.StatusInternalServerError)
	return
    }
//    areasWithCounter, err2 := fileAreaMapper.UpdateFileAreasWithFileCount(areas)
//    if err2 != nil {
//	response := fmt.Sprintf("Fail in UpdateFileAreasWithFileCount on fileAreaMapper: err = %+v", err2)
//	http.Error(w, response, http.StatusInternalServerError)
//	return
//    }
//    areasWithNewCounter, err3 := fileAreaMapper.UpdateNewFileAreasWithFileCount(areasWithCounter)
//    if err3 != nil {
//	response := fmt.Sprintf("Fail in UpdateNewFileAreasWithFileCount on fileAreaMapper: err = %+v", err3)
//	http.Error(w, response, http.StatusInternalServerError)
//	return
//    }

    // Шаг 2. Подготовка данных для вьюшки
    var areaHeaders []views.AreaHeader
    for _, a := range areas {
        areaHeaders = append(areaHeaders, views.AreaHeader{
            Name:        a.GetName(),
            Summary:     a.GetSummary(),
            IndexURL:    "/file/" + a.GetName(),
            NewMsgCount: a.GetNewCount(),  // новое имя поля для файлов
        })
    }
    data := views.FEchoAreaIndexData{
        Actions: []views.ToolbarAction{
            {Label: "Create", URL: "/file/create", Icon: "edit"},
        },
        Areas: areaHeaders,
    }
    err := views.Page("File Echoes", views.FEchoAreaIndexView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
