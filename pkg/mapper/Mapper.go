package mapper

import (
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
)

type Mapper struct {
	registry *registry.Container
}

func (self *Mapper) SetRegistry(r *registry.Container) {
	self.registry = r
}

func (self *Mapper) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}

func (self *Mapper) restoreMapperManager() *MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}