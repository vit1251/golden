package widgets

import (
	"io"
)

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
		menuAction.ID = "mainMenuHome"
		menuAction.Link = "/"
		menuAction.Label = "Home"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuDirect"
		menuAction.Link = "/netmail"
		menuAction.Label = "Netmail"
		menuAction.Metric = -1
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuEcho"
		menuAction.Link = "/echo"
		menuAction.Label = "Echomail"
		menuAction.Metric = -1
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuFile"
		menuAction.Link = "/file"
		menuAction.Label = "Files"
		menuAction.Metric = -1
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuService"
		menuAction.Link = "/service"
		menuAction.Label = "Service"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuStat"
		menuAction.Link = "/stat"
		menuAction.Label = "Statistics"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuTwit"
		menuAction.Link = "/twit"
		menuAction.Label = "Twit"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuDraft"
		menuAction.Link = "/draft"
		menuAction.Label = "Draft"
		menus = append(menus, menuAction)
	}

	if menuAction := NewMenuAction(); menuAction != nil {
		menuAction.ID = "mainMenuSetup"
		menuAction.Link = "/setup"
		menuAction.Label = "Setup"
		menus = append(menus, menuAction)
	}

	return menus
}

func (self *MainMenuWidget) Render(w io.Writer) error {
	self.mw.Render(w)
	return nil
}

func (self *MainMenuWidget) SetParam(ID string, value int) {
	//log.Printf("Set: name = %s value = %d", ID, value)
	for _, a := range self.mw.actions {
		//log.Printf("compare: %s == %s", a.ID, ID)
		if a.ID == ID {
			a.Metric = value
		}
	}
}
