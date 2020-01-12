package tosser

import (
	"path/filepath"
	"strings"
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
	var result bool = false
	var ext string = filepath.Ext(name)

	/* Monday packet */
	if strings.HasPrefix(ext, ".MO") {
		result = true
	}

	// ...

	/* Saturday packet */
	if strings.HasPrefix(ext, ".SA") {
		result = true
	}
	if strings.HasPrefix(ext, ".SU") {
		result = true
	}

	return result
}

func (self *Tosser) Toss() {
	self.ProcessInbound()
	self.ProcessOutbound()
}

func NewTosser() (*Tosser) {
	return new(Tosser)
}
