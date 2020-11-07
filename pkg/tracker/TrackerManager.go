package tracker

import (
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type TrackerManager struct {
	registry    *registry.Container
	event       chan bool
}

func NewTrackerManager(r *registry.Container) *TrackerManager {
	newTrackerManager := new(TrackerManager)
	newTrackerManager.registry = r
	newTrackerManager.event = make(chan bool)

	eventBus := newTrackerManager.restoreEventBus()
	eventBus.Register(newTrackerManager)

	return newTrackerManager
}

func (self *TrackerManager) HandleEvent(event string) {
	log.Printf("Tosser event receive")
	if event == "Tracker" {
		if self.event != nil {
			self.event <- true
		}
	}
}

func (self TrackerManager) Start() {
	go self.run()
}

func (self TrackerManager) processTosser() {
	newTosser := NewTracker(self.registry)
	newTosser.Track()
}

func (self *TrackerManager) run() {
	log.Printf(" * Tracker service start")
	var procIteration int
	for range self.event {
		log.Printf(" * Tracker start (%d)", procIteration)
		self.processTosser()
		log.Printf(" * Tracker complete (%d)", procIteration)
		procIteration += 1
	}
	log.Printf(" * Tracker service stop")
}

func (self TrackerManager) restoreEventBus() *eventbus.EventBus {
	ConfigManagerPtr := self.registry.Get("EventBus")
	if eventBus, ok := ConfigManagerPtr.(*eventbus.EventBus); ok {
		return eventBus
	} else {
		panic("no event bus")
	}
}