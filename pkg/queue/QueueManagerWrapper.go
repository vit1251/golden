package queue

import "github.com/vit1251/golden/pkg/registry"

const QUEUE_MANAGER_ID = "QueueManager"

func RestoreQueueManager(r *registry.Container) *QueueManager {
	managerPtr := r.Get(QUEUE_MANAGER_ID)
	if manager, ok := managerPtr.(*QueueManager); ok {
		return manager
	} else {
		panic("no queue manager")
	}
}
