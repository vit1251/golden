package action

import (
	"fmt"
	"net/http"
)

type ServiceTrackerEventAction struct {
	Action
}

func NewServiceTrackerEventAction() *ServiceTrackerEventAction {
	return new(ServiceTrackerEventAction)
}

func (self *ServiceTrackerEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := self.restoreEventBus()

	/* Create mailer event */
	newMailerEvent := eventBus.CreateEvent("Tracker")
	eventBus.FireEvent(newMailerEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service/tracker/stat")
	http.Redirect(w, r, newLocation, 303)

}
