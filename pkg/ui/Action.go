package ui

import (
	"net/http"
)

type Action struct {
	Site *WebSite
}

func (self *Action) SetSite(webSite *WebSite) {
	self.Site = webSite
}

func (self *Action) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (self *Action) GetHandler() func(http.ResponseWriter, *http.Request) {
	return self.ServeHTTP
}
