package action

import (
	"fmt"
	"net/http"
)

type ServiceTossEventAction struct {
	Action
}

func NewServiceTossEventAction() *ServiceTossEventAction {
	return new(ServiceTossEventAction)
}

func (self *ServiceTossEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := self.restoreEventBus()

	/* Create mailer event */
	newMailerEvent := eventBus.CreateEvent("Tosser")
	eventBus.FireEvent(newMailerEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service/toss/stat")
	http.Redirect(w, r, newLocation, 303)

}
