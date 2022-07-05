package config

import "github.com/vit1251/golden/pkg/registry"

const CONFIG_MANAGER_ID = "ConfigManager"

func RestoreConfigManager(r *registry.Container) *ConfigManager {
	managerPtr := r.Get(CONFIG_MANAGER_ID)
	if manager, ok := managerPtr.(*ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}
