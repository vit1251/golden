package tosser

import (
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"time"
)

type Tosser struct {
	registry *registry.Container
}

func NewTosser(registry *registry.Container) *Tosser {
	tosser := new(Tosser)
	tosser.registry = registry
	return tosser
}

func (self *Tosser) Toss() {

	tosserStart := time.Now()
	log.Printf("Start tosser session")

	err1 := self.ProcessInbound()
	if err1 != nil {
		log.Printf("err = %+v", err1)
	}
	err2 := self.ProcessOutbound()
	if err2 != nil {
		log.Printf("err = %+v", err2)
	}

	log.Printf("Stop tosser session")
	elapsed := time.Since(tosserStart)
	log.Printf("Tosser session: %+v", elapsed)
}

func (self Tosser) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}

func (self Tosser) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}

func (self Tosser) restoreMapperManager() *mapper.MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*mapper.MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}

func (self *Tosser) restoreConfigManager() *config.ConfigManager {
	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*config.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}
