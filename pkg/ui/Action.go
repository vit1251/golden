package ui

import (
)

type Action struct {
	Site *WebSite
}

func (self *Action) SetSite(webSite *WebSite) {
	self.Site = webSite
}
