package tosser

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

func (self *Tosser) Toss() {
	self.ProcessInbound()
	self.ProcessOutbound()
}

func NewTosser() (*Tosser) {
	return new(Tosser)
}
