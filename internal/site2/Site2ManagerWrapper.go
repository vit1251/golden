package site2

import (
	"github.com/vit1251/golden/pkg/registry"
)

const SITE2_MANAGER_ID = "Site2Manager"

func RestoreSite2Manager(r *registry.Container) *Site2Manager {
	site2ManagerPtr := r.Get(SITE2_MANAGER_ID)
	if site2Manager, ok := site2ManagerPtr.(*Site2Manager); ok {
		return site2Manager
	} else {
		panic("no site2 manager")
	}
}
