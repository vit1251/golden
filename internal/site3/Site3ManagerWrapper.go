package site3

import (
	"github.com/vit1251/golden/pkg/registry"
)

const SITE3_MANAGER_ID = "Site3Manager"

func RestoreSite3Manager(r *registry.Container) *Site3Manager {
	site3ManagerPtr := r.Get(SITE3_MANAGER_ID)
	if site3Manager, ok := site3ManagerPtr.(*Site3Manager); ok {
		return site3Manager
	} else {
		panic("no site3 manager")
	}
}
