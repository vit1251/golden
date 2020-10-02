package action

import (
	"go.uber.org/dig"
	"net/http"
)

type IAction interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Action struct {
	IAction
	Container *dig.Container
}

func (self *Action) SetContainer(c *dig.Container) {
	self.Container = c
}
