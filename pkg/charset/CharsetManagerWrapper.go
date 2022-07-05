package charset

import "github.com/vit1251/golden/pkg/registry"

const CHARSET_MANAGER_ID = "CharsetManager"

func RestoreCharsetManager(r *registry.Container) *CharsetManager {
	managerPtr := r.Get(CHARSET_MANAGER_ID)
	if manager, ok := managerPtr.(*CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}
