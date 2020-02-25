package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/common"
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

	master := common.GetMaster()
	fileManager := master.FileManager

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)
	file := vars["file"]
	log.Printf("file = %v", file)

	/* Get message area */
	area, err1 := fileManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on FileManager")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	/* Path */
	path := filepath.Join(area.Path, file)

	/* Open original file */
	stream, err2 := os.Open(path)
	if err2 != nil {
		response := fmt.Sprintf("Fail on Open")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	defer stream.Close()

	/* Transmit original file */
	io.Copy(w, stream)

}