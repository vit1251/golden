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

type Event struct {
	name string
}

func NewEventBus(registry *registry.Container) *EventBus {
	bus := new(EventBus)
	return bus
}

func (self *EventBus) CreateEvent(eventName string) Event {
	return Event{
		name: eventName,
	}
}

func (self *EventBus) FireEvent(evt Event) {
	log.Printf("EventBus: event = %#v", evt)
	for _, action := range self.handlers {
		log.Printf("EventBus: evt = %#v action = %q", evt, action)
		action.HandleEvent(evt.name)
	}
}

func (self *EventBus) Register(action IEventBusAction) {
	self.handlers = append(self.handlers, action)
}
