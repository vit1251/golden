package eventbus

import (
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type IEventBusAction interface {
	HandleEvent(event string)
}

type EventBus struct {
	handlers []IEventBusAction
}

func NewEventBus(registry *registry.Container) *EventBus {
	bus := new(EventBus)
	return bus
}

func (self *EventBus) Event(event string) {
	log.Printf("EventBus: event = %s", event)
	for _, action := range self.handlers {
		log.Printf("EventBus: event = %s action = %q", event, action)
		action.HandleEvent(event)
	}
}

func (self *EventBus) Register(action IEventBusAction) {
	self.handlers = append(self.handlers, action)
}
