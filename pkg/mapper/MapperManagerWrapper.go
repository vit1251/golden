package mapper

import "github.com/vit1251/golden/pkg/registry"

const MAPPER_MANAGER_ID = "MapperManager"

func RestoreMapperManager(r *registry.Container) *MapperManager {
	managerPtr := r.Get(MAPPER_MANAGER_ID)
	if manager, ok := managerPtr.(*MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}
