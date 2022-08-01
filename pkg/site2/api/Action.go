package api

import (
	"github.com/vit1251/golden/pkg/registry"
	"net/http"
)

type Handler func(req []byte) []byte

type IAction interface {
	SetContainer(r *registry.Container)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Action struct {
	Type   string
	r      *registry.Container
	Handle Handler
}

func (self *Action) SetContainer(r *registry.Container) {
	self.r = r
}

func (self *Action) SetRegistry(r *registry.Container) {
	self.r = r
}

func (self *Action) GetRegistry() *registry.Container {
	return self.r
}
