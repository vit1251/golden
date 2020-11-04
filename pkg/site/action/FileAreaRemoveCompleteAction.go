package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type FileAreaRemoveCompleteAction struct {
	Action
}

func NewFileAreaRemoveCompleteAction() *FileAreaRemoveCompleteAction {
	return new(FileAreaRemoveCompleteAction)
}

func (self FileAreaRemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fileManager := self.restoreFileManager()

	/* Restore area name */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Restore area by name */
	area, err1 := fileManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Restore areas file */
	items, err2 := fileManager.GetFileHeaders(echoTag)
	if err2 != nil {
		panic(err2)
	}

	/* Remove files */
	var areaName string = area.GetName()
	for _, i := range items {
		newPath := fileManager.GetFileAbsolutePath(areaName, i.File)
		log.Printf("Remove(%s)", newPath)
		//err1 := os.Remove(newPath)
		//if err1 != nil {
		//	log.Printf("Fail on Remove: name = %+v err = %+v", newPath, err1)
		//	panic(err1)
		//}
	}

	/* Remove directory with box */
	newBoxPath := fileManager.GetFileBoxAbsolutePath(areaName)
	log.Printf("RemoveAll(%s)", newBoxPath)
	err3 := os.RemoveAll(newBoxPath)
	if err3 != nil {
		log.Printf("Fail on RemoveAll(%s): err = %+v", newBoxPath, err3)
	}

	/* Remove files in area */
	err4 := fileManager.RemoveFilesByAreaName(echoTag)
	if err4 != nil {
		log.Printf("err4 = %+v", err4)
	}

	/* Remove area by name */
	err5 := fileManager.RemoveAreaByName(echoTag)
	if err5 != nil {
		log.Printf("err5 = %+v", err5)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/file")
	http.Redirect(w, r, newLocation, 303)

}