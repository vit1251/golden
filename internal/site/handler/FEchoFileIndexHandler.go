package handler

import (
    "net/http"
    "os"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoFileIndexHandler struct {
    registry *registry.Container
}

func NewFEchoFileIndexHandler(registry *registry.Container) *FEchoFileIndexHandler {
    return &FEchoFileIndexHandler{
    	registry: registry,
    }
}

func (h *FEchoFileIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()
    fileMapper := mapperManager.GetFileMapper()

    areaIndex := r.PathValue("echoname")
    area, err1 := fileAreaMapper.GetAreaByName(areaIndex)
    if err1 != nil || area == nil {
	http.Error(w, "Area not found", http.StatusInternalServerError)
	return
    }

    files, err2 := fileMapper.GetFileHeaders(areaIndex)
    if err2 != nil {
	http.Error(w, err2.Error(), http.StatusInternalServerError)
	return
    }

    var rows []views.FEchoFileRow
    for _, f := range files {
        path := fileMapper.GetFileAbsolutePath(areaIndex, f.GetFile())
        _, statErr := os.Stat(path)

        rows = append(rows, views.FEchoFileRow{
            OrigName:  f.GetOrigName(),
            Desc:      f.GetDesc(),
            Date:      f.GetTime().Format("2006-01-02 15:04"),
            ViewURL:   "/file/" + areaIndex + "/tic/" + f.GetFile() + "/view",
            IsNew:     f.GetViewCount() == 0,
            IsMissing: os.IsNotExist(statErr),
        })
    }

    data := views.FEchoFileIndexData{
        Actions: []views.ToolbarAction{
            {Label: "Upload",   URL: "/file/" + areaIndex + "/upload", Icon: "edit"},
            {Label: "Settings", URL: "/file/" + areaIndex + "/update", Icon: "settings"},
        },
        AreaName: area.GetName(),
        Files:    rows,
    }

    err := views.Page(area.GetName(), views.FEchoFileIndexView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

