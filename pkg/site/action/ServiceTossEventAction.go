package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/eventbus"
	"net/http"
)

type ServiceTossEventAction struct {
	Action
}

func NewServiceTossEventAction() *ServiceTossEventAction {
	return new(ServiceTossEventAction)
}

func (self *ServiceTossEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := eventbus.RestoreEventBus(self.GetRegistry())

	/* Create mailer event */
	newMailerEvent := eventBus.CreateEvent("Tosser")
	eventBus.FireEvent(newMailerEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service/toss/stat")
	http.Redirect(w, r, newLocation, 303)

}
