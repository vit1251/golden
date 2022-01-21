package action

import (
	"fmt"
	"net/http"
)

type ServiceMailerEventAction struct {
	Action
}

func NewServiceMailerEventAction() *ServiceMailerEventAction {
	return new(ServiceMailerEventAction)
}

func (self *ServiceMailerEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := self.restoreEventBus()

	/* Create mailer event */
	newMailerEvent := eventBus.CreateEvent("Mailer")
	eventBus.FireEvent(newMailerEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service/mailer/stat")
	http.Redirect(w, r, newLocation, 303)

}
