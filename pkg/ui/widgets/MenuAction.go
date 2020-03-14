package widgets

type MenuAction struct {
	Link  string
	Icon  string
	Label string
}

func NewMenuAction() *MenuAction {
	ma := new(MenuAction)
	return ma
}
