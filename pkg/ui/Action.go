package ui

import (
)

type Action struct {
	Site *GoldenSite
}

func (self *Action) SetSite(webSite *GoldenSite) {
	self.Site = webSite
}
