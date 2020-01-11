package tosser

import (
	"path/filepath"
)

type Tosser struct {
	inboundDirectory      string
	workInboundDirectory  string
}

func (self *Tosser) SetInboundDirectory(inboundDirectory string) {
	self.inboundDirectory = inboundDirectory
}

func (self *Tosser) SetWorkInboundDirectory(workInboundDirectory string) {
	self.workInboundDirectory = workInboundDirectory
}

func IsNetmail(name string) bool {
	var ext string = filepath.Ext(name)
	return ext == ".pkt"
}

func IsArchmail(name string) bool {
	var ext string = filepath.Ext(name)
	return ext != ".pkt"
}

func (self *Tosser) Toss() {
	self.ProcessInbound()
	self.ProcessOutbound()
}

func NewTosser() (*Tosser) {
	return new(Tosser)
}
