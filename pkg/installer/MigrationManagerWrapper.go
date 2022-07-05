package installer

import "github.com/vit1251/golden/pkg/registry"

const MIGRATION_MANAGER_ID = "MigrationManager"

func RestoreMigrationManager(r *registry.Container) *MigrationManager {
	managerPtr := r.Get(MIGRATION_MANAGER_ID)
	if manager, ok := managerPtr.(*MigrationManager); ok {
		return manager
	} else {
		panic("no migration manager")
	}
}
