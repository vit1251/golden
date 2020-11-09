package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type FileEchoAreaRemoveAction struct {
	Action
}

func NewFileEchoAreaRemoveAction() *FileEchoAreaRemoveAction {
	fa := new(FileEchoAreaRemoveAction)
	return fa
}


func (self *FileEchoAreaRemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
