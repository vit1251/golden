package action

import (
	"fmt"
	"net/http"
)

type FileAreaComposeCompleteAction struct {
	Action
}

func NewFileAreaComposeCompleteAction() *FileAreaComposeAction {
	return new(FileAreaComposeAction)
}

func (self *FileAreaComposeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//fileArea := "NASA"

	// TODO - ...

	/* Redirect */
	newLocation := fmt.Sprintf("/file")
	http.Redirect(w, r, newLocation, 303)

}
