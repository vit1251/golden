package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/setup"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type FileAreaDownloadAction struct {
	Action
}

func NewFileAreaDownloadAction() *FileAreaDownloadAction {
	fa := new(FileAreaDownloadAction)
	return fa
}

func (self *FileAreaDownloadAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var setupManager *setup.ConfigManager
	var fileManager *file.FileManager
	self.Container.Invoke(func(fm *file.FileManager, sm *setup.ConfigManager) {
		fileManager = fm
		setupManager = sm
	})

	//
	storagePath, err1 := setupManager.Get("main", "FileBox", "Files")
	if err1 != nil {
		response := fmt.Sprintf("Fail on get parameters")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)
	newFile := vars["file"]
	log.Printf("file = %v", newFile)

	/* Get message area */
	area, err1 := fileManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on FileManager")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	//fileManager.GetFileByName(newFile)

	/* Path */
	var areaName string = area.Name()
	path := filepath.Join(storagePath, areaName, newFile)

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