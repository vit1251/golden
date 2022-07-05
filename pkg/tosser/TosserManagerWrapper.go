package tosser

import "github.com/vit1251/golden/pkg/registry"

const TOSSER_MANAGER_ID = "TosserManager"

func RestoreTosserManager(r *registry.Container) *TosserManager {
	managerPtr := r.Get(TOSSER_MANAGER_ID)
	if manager, ok := managerPtr.(*TosserManager); ok {
		return manager
	} else {
		panic("no tosser manager")
	}
}
