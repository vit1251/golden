package handler

import (
    "archive/zip"
    "net/http"
    "os"
    "strings"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
    "github.com/vit1251/golden/internal/utils"
)

type FEchoFileViewHandler struct {
    registry *registry.Container
}

func NewFEchoFileViewHandler(registry *registry.Container) *FEchoFileViewHandler {
    return &FEchoFileViewHandler{
	registry: registry,
    }
}

func (h *FEchoFileViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()
    fileMapper := mapperManager.GetFileMapper()

    areaName := r.PathValue("echoname")
    indexName := r.PathValue("file")

    area, err := fileAreaMapper.GetAreaByName(areaName)
    if err != nil || area == nil {
        http.Error(w, "Area not found", http.StatusInternalServerError)
        return
    }

    file, err := fileMapper.GetFileByIndexName(areaName, indexName)
    if err != nil || file == nil {
        http.Error(w, "File not found", http.StatusInternalServerError)
        return
    }

    fileMapper.ViewFileByIndexName(areaName, indexName)

    origPath := file.GetAbsolutePath()
    origName := file.GetOrigName()

    actions := []views.ToolbarAction{
        {Label: "Back",     URL: "/file/" + areaName,                                     Icon: "arrow-left"},
        {Label: "Download", URL: "/file/" + areaName + "/tic/" + indexName + "/download", Icon: "edit"},
        {Label: "Archive",  URL: "/file/" + areaName + "/tic/" + indexName + "/archive",  Icon: "archive"},
    }

    data := views.FEchoFileViewData{
        Actions:  actions,
        OrigName: origName,
        Desc:     file.GetDesc(),
        Origin:   file.GetOrigin(),
        From:     file.GetFrom(),
        To:       file.GetTo(),
        DiskPath: origPath,
        DiskSize: "?",
        Crc:      file.GetCrc(),
        ImageURL: "",
    }

    fi, err := os.Stat(origPath)
    if err == nil {
        data.DiskSize = utils.FormatBytes(fi.Size())
    }

    if IsImage(origName) {
        data.ImageURL = "/file/" + areaName + "/tic/" + indexName + "/download"
    } else if IsZipArchive(origName) {
        reader, err := zip.OpenReader(origPath)
        if err == nil {
            data.ZipComment = strings.Replace(reader.Comment, "\r\n", "\n", -1)
            for _, f := range reader.File {
                var uncompressedSize int64 = int64(f.UncompressedSize64)
                data.ZipFiles = append(data.ZipFiles, views.ZipEntry{
                    Name:    f.Name,
                    Comment: f.Comment,
                    Size:    utils.FormatBytes(uncompressedSize),
                })
            }
            reader.Close()
        }
    }

    err = views.Page(area.GetName(), views.FEchoFileViewView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func IsImage(filename string) bool {
    upper := strings.ToUpper(filename)
    return strings.HasSuffix(upper, ".GIF") || strings.HasSuffix(upper, ".JPG") || strings.HasSuffix(upper, ".PNG")
}

func IsZipArchive(filename string) bool {
    return strings.HasSuffix(strings.ToUpper(filename), ".ZIP")
}
