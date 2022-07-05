package tracker

import "github.com/vit1251/golden/pkg/registry"

const TRACKER_MANAGER_ID = "TrackerManager"

func RestoreTrackerManager(r *registry.Container) *TrackerManager {
	managerPtr := r.Get(TRACKER_MANAGER_ID)
	if manager, ok := managerPtr.(*TrackerManager); ok {
		return manager
	} else {
		panic("no tracker manager")
	}
}
