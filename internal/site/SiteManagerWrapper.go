package site

import (
	"github.com/vit1251/golden/pkg/registry"
)

const SITE_MANAGER_ID = "SiteManager"

func RestoreSiteManager(r *registry.Container) *SiteManager {
	siteManagerPtr := r.Get(SITE_MANAGER_ID)
	if siteManager, ok := siteManagerPtr.(*SiteManager); ok {
		return siteManager
	} else {
		panic("no site manager")
	}
}
