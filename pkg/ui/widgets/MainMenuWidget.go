package widgets

import "net/http"

type MainMenuWidget struct {
	mw *MenuWidget
}

func NewMainMenuWidget() *MainMenuWidget {
	mmw := new(MainMenuWidget)
	mmw.mw = NewMenuWidget()
	mmw.mw.actions = mmw.Init()
	return mmw
}

func (self *MainMenuWidget) Init() []*MenuAction {

	var menus []*MenuAction

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/"
		menuAction.Label = "Home"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/netmail"
		menuAction.Label = "Netmail"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/echo"
		menuAction.Label = "Echomail"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/file"
		menuAction.Label = "Filebox"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/stat"
		menuAction.Label = "Statistics"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/service"
		menuAction.Label = "Service"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.Link = "/setup"
		menuAction.Label = "Setup"
		menus = append(menus, menuAction)
	}

	return menus
}

func (self *MainMenuWidget) Render(w http.ResponseWriter) error {
	self.mw.Render(w)
	return nil
}
