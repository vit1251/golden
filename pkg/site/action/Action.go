package action

import (
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type IAction interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Action struct {
	IAction
	registry *registry.Container
}

func (self *Action) GetRegistry() *registry.Container {
	return self.registry
}

func (self *Action) SetContainer(r *registry.Container) {
	self.registry = r
}

func (self Action) makeMenu() *widgets.MainMenuWidget {

	mapperManager := mapper.RestoreMapperManager(self.registry)
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
