package um

import "github.com/vit1251/golden/pkg/registry"

const URL_MANAGER_ID = "UrlManager"

func RestoreUrlManager(r *registry.Container) *UrlManager {
	managerPtr := r.Get(URL_MANAGER_ID)
	if manager, ok := managerPtr.(*UrlManager); ok {
		return manager
	} else {
		panic("no url manager")
	}
}
