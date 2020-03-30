package msg

import "strings"

type Area struct {
	name        string
	Count       int
	MsgNewCount int
}

func NewArea() *Area {
	a := new(Area)
	return a
}

func (self *Area) Name() string {
	return self.name
}

func (self *Area) SetName(name string) {
	self.name = strings.ToUpper(name)
}