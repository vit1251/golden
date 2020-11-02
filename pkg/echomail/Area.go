package echomail

import "strings"

type Area struct {
	name            string     /* Echo name              */
	Summary         string     /* Echo summary           */
	Charset         string     /* Echo charset           */
	MessageCount    int        /* Echo message count     */
	NewMessageCount int        /* Echo new message count */
}

func NewArea() *Area {
	a := new(Area)
	a.Charset = "CP866"
	return a
}

func (self *Area) GetName() string {
	return self.name
}

func (self *Area) SetName(name string) {
	self.name = strings.ToUpper(name)
}
