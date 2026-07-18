package handler

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoFileDownloadHandler struct {
    registry *registry.Container
}

func NewFEchoFileDownloadHandler(registry *registry.Container) *FEchoFileDownloadHandler {
    return &FEchoFileDownloadHandler{
	registry: registry,
    }
}

func (h *FEchoFileDownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()
    fileMapper := mapperManager.GetFileMapper()

	/* Parse URL parameters */
	var areaIndex string = r.PathValue("echoname")
	log.Printf("echoTag = %v", areaIndex)

	var newFile string = r.PathValue("file")
	log.Printf("file = %v", newFile)

	/* Get message area */
	area, err1 := fileAreaMapper.GetAreaByName(areaIndex)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	/* Path */
	var areaName string = area.GetName()
	path := fileMapper.GetFileAbsolutePath(areaName, newFile)

	/* Open original file */
	stream, err2 := os.Open(path)
	if err2 != nil {
		response := fmt.Sprintf("Fail on open source %s", path)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	defer stream.Close()

	/* Transmit original file */
	var sourceName string = fmt.Sprintf("attachment; filename=\"%s\"", newFile)
	w.Header().Set("Content-Disposition", sourceName)

    io.Copy(w, stream)

}
