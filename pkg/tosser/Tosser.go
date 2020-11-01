package tosser

import (
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/echomail"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"log"
	"time"
)

type Tosser struct {
	registry       *registry.Container
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

func (self *Tosser) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}

func (self *Tosser) restoreNetmailManager() *netmail.NetmailManager {
	managerPtr := self.registry.Get("NetmailManager")
	if manager, ok := managerPtr.(*netmail.NetmailManager); ok {
		return manager
	} else {
		panic("no netmail manager")
	}
}

func (self *Tosser) restoreAreaManager() *echomail.AreaManager {
	managerPtr := self.registry.Get("AreaManager")
	if manager, ok := managerPtr.(*echomail.AreaManager); ok {
		return manager
	} else {
		panic("no area manager")
	}
}

func (self *Tosser) restoreMessageManager() *echomail.MessageManager {
	managerPtr := self.registry.Get("MessageManager")
	if manager, ok := managerPtr.(*echomail.MessageManager); ok {
		return manager
	} else {
		panic("no message manager")
	}
}

func (self *Tosser) restoreStatManager() *stat.StatManager {
	managerPtr := self.registry.Get("StatManager")
	if manager, ok := managerPtr.(*stat.StatManager); ok {
		return manager
	} else {
		panic("no stat manager")
	}
}

func (self *Tosser) restoreConfigManager() *setup.ConfigManager {
	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*setup.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}

func (self *Tosser) restoreFileManager() *file.FileManager {
	managerPtr := self.registry.Get("FileManager")
	if manager, ok := managerPtr.(*file.FileManager); ok {
		return manager
	} else {
		panic("no file manager")
	}
}

