package action

import (
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/site/widgets"
	"github.com/vit1251/golden/pkg/storage"
	"github.com/vit1251/golden/pkg/tosser"
	"github.com/vit1251/golden/pkg/um"
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

func (self Action) makeMenu() *widgets.MainMenuWidget {

	mapperManager := self.restoreMapperManager()
	echoMapper := mapperManager.GetEchoMapper()
	netmailMapper := mapperManager.GetNetmailMapper()
	fileMapper := mapperManager.GetFileMapper()

	newCount, _ := echoMapper.GetMessageNewCount()
	newDirect, _ := netmailMapper.GetMessageNewCount()
	newFile, _ := fileMapper.GetFileNewCount()

	mainMenu := widgets.NewMainMenuWidget()
	mainMenu.SetParam("mainMenuEcho", newCount)
	mainMenu.SetParam("mainMenuDirect", newDirect)
	mainMenu.SetParam("mainMenuFile", newFile)

	return mainMenu
}

func (self Action) restoreTosserManager() *tosser.TosserManager {
	managerPtr := self.registry.Get("TosserManager")
	if manager, ok := managerPtr.(*tosser.TosserManager); ok {
		return manager
	} else {
		panic("no tosser manager")
	}
}

func (self Action) restoreEventBus() *eventbus.EventBus {
	managerPtr := self.registry.Get("EventBus")
	if manager, ok := managerPtr.(*eventbus.EventBus); ok {
		return manager
	} else {
		panic("no eventbus manager")
	}
}

func (self Action) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}

func (self Action) restoreMapperManager() *mapper.MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*mapper.MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}

func (self Action) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}

func (self *Action) restoreConfigManager() *config.ConfigManager {
	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*config.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}

func (self *Action) restoreUrlManager() *um.UrlManager {
	managerPtr := self.registry.Get("UrlManager")
	if manager, ok := managerPtr.(*um.UrlManager); ok {
		return manager
	} else {
		panic("no url manager")
	}
}
