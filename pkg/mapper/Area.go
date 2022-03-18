package mapper

import (
	"log"
	"strings"
)

type Area struct {
	name            string /* Echo name              */
	Summary         string /* Echo summary           */
	charset         string /* Echo charset           */
	MessageCount    int    /* Echo message count     */
	newMessageCount int    /* Echo new message count */
	order           int64  /* Area sort order        */
}

func (self *Area) GetNewMessageCount() int {
	return self.newMessageCount
}

func (self *Area) SetNewMessageCount(newMessageCount int) {
	self.newMessageCount = newMessageCount
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
		log.Printf("Warning: no charset for %s. Set 'CP866' as default.", self.name)
		self.charset = "CP866"
	}
	return self.charset
}

func (self *Area) GetOrder() int64 {
	return self.order
}

func (self *Area) SetOrder(order int64) {
	self.order = order
}

func (self *Area) GetSummary() string {
	return self.Summary
}

func (self *Area) SetSummary(summary string) {
	self.Summary = summary
}
