package mapper

import (
	"github.com/vit1251/golden/pkg/registry"
)

type Mapper struct {
	registry *registry.Container
}

func (self *Mapper) SetRegistry(r *registry.Container) {
	self.registry = r
}
