package storage

import "github.com/vit1251/golden/pkg/registry"

const STORAGE_MANAGER_ID = "StorageManager"

func RestoreStorageManager(r *registry.Container) *StorageManager {
	managerPtr := r.Get(STORAGE_MANAGER_ID)
	if manager, ok := managerPtr.(*StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}
