package handler

import (
	"fmt"
	"net/http"

	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/registry"
)

type ServiceTrackerEventHandler struct {
	registry *registry.Container
}

func NewServiceTrackerEventHandler(registry *registry.Container) *ServiceTrackerEventHandler {
	return &ServiceTrackerEventHandler{
		registry: registry,
	}
}

func (self *ServiceTrackerEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := eventbus.RestoreEventBus(self.registry)

	/* Create mailer event */
	newMailerEvent := eventBus.CreateEvent("Tracker")
	eventBus.FireEvent(newMailerEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service/tracker/stat")
	http.Redirect(w, r, newLocation, 303)

}
