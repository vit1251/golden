package eventbus

import "github.com/vit1251/golden/pkg/registry"

const EVENT_BUS_MANAGER_ID = "EventBus"

func RestoreEventBus(r *registry.Container) *EventBus {
	managerPtr := r.Get(EVENT_BUS_MANAGER_ID)
	if manager, ok := managerPtr.(*EventBus); ok {
		return manager
	} else {
		panic("no eventbus manager")
	}
}
