package widgets

type UserAction struct {
	Link  string
	Icon  string
	Label string
}

func NewUserAction() *UserAction {
	ua:=new(UserAction)
	return ua
}