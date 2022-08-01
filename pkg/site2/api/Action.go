package api

import (
	"github.com/vit1251/golden/pkg/registry"
)

type Handler func() []byte

type Action struct {
	Type    string
	r      *registry.Container
	Handle  Handler
}

func (self *Action) SetRegistry(r *registry.Container) {
	self.r = r
}

func (self *Action) GetRegistry() *registry.Container {
	return self.r
}
