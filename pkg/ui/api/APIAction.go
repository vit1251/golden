package api

import (
	"go.uber.org/dig"
)

type APIAction struct {
	Container *dig.Container
}

func (self *APIAction) SetContainer(c *dig.Container) {
	self.Container = c
}
