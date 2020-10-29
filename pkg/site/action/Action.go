package action

import (
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/tosser"
	"net/http"
)

type IAction interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Action struct {
	IAction
	registry *registry.Container
}

func (self *Action) SetContainer(r *registry.Container) {
	self.registry = r
}

func (self *Action) restoreMessageManager() *msg.MessageManager {

	managerPtr := self.registry.Get("MessageManager")
	if manager, ok := managerPtr.(*msg.MessageManager); ok {
		return manager
	} else {
		panic("no message manager")
	}

}

func (self *Action) restoreAreaManager() *msg.AreaManager {

	managerPtr := self.registry.Get("AreaManager")
	if manager, ok := managerPtr.(*msg.AreaManager); ok {
		return manager
	} else {
		panic("no area manager")
	}

}

func (self *Action) restoreTosserManager() *tosser.TosserManager {
	managerPtr := self.registry.Get("TosserManager")
	if manager, ok := managerPtr.(*tosser.TosserManager); ok {
		return manager
	} else {
		panic("no tosser manager")
	}
}

func (self *Action) restoreStatManager() *stat.StatManager {
	managerPtr := self.registry.Get("StatManager")
	if manager, ok := managerPtr.(*stat.StatManager); ok {
		return manager
	} else {
		panic("no stat manager")
	}
}

func (self *Action) restoreConfigManager() *setup.ConfigManager {
	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*setup.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}

func (self *Action) restoreFileManager() *file.FileManager {
	managerPtr := self.registry.Get("FileManager")
	if manager, ok := managerPtr.(*file.FileManager); ok {
		return manager
	} else {
		panic("no filemanager manager")
	}
}

func (self *Action) restoreNetmailManager() *netmail.NetmailManager {
	managerPtr := self.registry.Get("NetmailManager")
	if manager, ok := managerPtr.(*netmail.NetmailManager); ok {
		return manager
	} else {
		panic("no netmail manager")
	}
}
