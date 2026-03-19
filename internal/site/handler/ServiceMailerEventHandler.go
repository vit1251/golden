package handler

import (
	"fmt"
	"net/http"

	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/registry"
)

type ServiceMailerEventHandler struct {
	registry *registry.Container
}

func NewServiceMailerEventHandler(registry *registry.Container) *ServiceMailerEventHandler {
	return &ServiceMailerEventHandler{
		registry: registry,
	}
}

func (self *ServiceMailerEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eventBus := eventbus.RestoreEventBus(self.registry)

	/* Create mailer event */
	newMailerEvent := eventBus.CreateEvent("Mailer")
	eventBus.FireEvent(newMailerEvent)

	/* Redirect */
	newLocation := fmt.Sprintf("/service/mailer/stat")
	http.Redirect(w, r, newLocation, 303)

}
