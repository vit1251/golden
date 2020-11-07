package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

type FileAreaDownloadAction struct {
	Action
}

func NewFileAreaDownloadAction() *FileAreaDownloadAction {
	fa := new(FileAreaDownloadAction)
	return fa
}

func (self *FileAreaDownloadAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	fileMapper := mapperManager.GetFileMapper()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)
	newFile := vars["file"]
	log.Printf("file = %v", newFile)

	/* Get message area */
	area, err1 := fileMapper.GetAreaByName(echoTag)
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
