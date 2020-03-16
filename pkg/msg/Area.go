package msg

import "strings"

type Area struct {
	Name        string
	Count       int
	MsgNewCount int
}

func NewArea() *Area {
	a := new(Area)
	return a
}

func (self *Area) SetName(name string) {
	self.Name = strings.ToUpper(name)
}