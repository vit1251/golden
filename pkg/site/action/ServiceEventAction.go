package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type ServiceEventAction struct {
	Action
}

func NewServiceEventAction() *ServiceEventAction {
	return new(ServiceEventAction)
}

func (self *ServiceEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := self.restoreEventBus()

	vars := mux.Vars(r)

	/* Recover message */
	servceEvent := vars["name"]

	newEvent := strings.Title(servceEvent)

	eventBus.Event(newEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service")
	http.Redirect(w, r, newLocation, 303)

}
