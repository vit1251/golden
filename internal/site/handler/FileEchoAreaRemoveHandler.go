package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type FileEchoAreaRemoveHandler struct {
	registry *registry.Container
}

func NewFileEchoAreaRemoveHandler(registry *registry.Container) *FileEchoAreaRemoveHandler {
	return &FileEchoAreaRemoveHandler{
		registry: registry,
	}
}

func (self *FileEchoAreaRemoveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
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

	/* Remove records */
	err4 := fileMapper.RemoveFileByName(newFile)
	if err4 != nil {
		response := fmt.Sprintf("Fail on RemoveFileByName in fileMapper: err = %+v", err4)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Remove attachment */
	var areaName string = area.GetName()
	newPath := fileMapper.GetFileAbsolutePath(areaName, newFile)
	err5 := os.Remove(newPath)
	if err5 != nil {
		log.Printf("Fail remove %s: err = %+v", newPath, err5)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/file/%s", area.GetName())
	http.Redirect(w, r, newLocation, 303)

}
