package ui

import (
	"go.uber.org/dig"
)

type Action struct {
	Container *dig.Container
}

func (self *Action) SetContainer(c *dig.Container) {
	self.Container = c
}
