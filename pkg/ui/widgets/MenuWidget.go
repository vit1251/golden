package widgets

type MenuWidget struct {
	actions []*MenuAction
}

func (self* MenuWidget) Add(manuAction *MenuAction) {
	self.actions = append(self.actions, manuAction)
}

func (self* MenuWidget) Actions() []*MenuAction {
	return self.actions
}

func NewMenuWidget() *MenuWidget {
	mw := new(MenuWidget)
	return mw
}

