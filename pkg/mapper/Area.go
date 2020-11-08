package mapper

import (
	"log"
	"strings"
)

type Area struct {
	name            string     /* Echo name              */
	Summary         string     /* Echo summary           */
	charset         string     /* Echo charset           */
	MessageCount    int        /* Echo message count     */
	NewMessageCount int        /* Echo new message count */
}

func NewArea() *Area {
	a := new(Area)
	a.charset = "CP866"
	return a
}

func (self *Area) GetName() string {
	return self.name
}

func (self *Area) SetName(name string) {
	self.name = strings.ToUpper(name)
}

func (self *Area) SetCharset(newCharset string) {
	if newCharset == "CP866" || newCharset == "UTF-8" {
		self.charset = newCharset
	}
}

func (self Area) GetCharset() string {
	if self.charset == "" {
		log.Printf("Area: Warning no charset pre-define use CP866.")
		self.charset = "CP866"
	}
	return self.charset
}
