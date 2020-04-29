package widgets

type MenuAction struct {
	Link   string
	Icon   string
	Label  string
	Metric int
	ID     string
}

func (self *MenuAction) SetLink(s string) *MenuAction {
	self.Link = s
	return self
}

func (self *MenuAction) SetIcon(s string) *MenuAction {
	self.Icon = s
	return self
}

func (self *MenuAction) SetLabel(s string) *MenuAction {
	self.Label = s
	return self
}

func NewMenuAction() *MenuAction {
	ma := new(MenuAction)
	return ma
}
