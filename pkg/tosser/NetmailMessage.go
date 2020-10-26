package tosser

type NetmailMessage struct {
	Subject string
	To      string
	ToAddr  string
	From    string
	body    string
	Reply   string
	Kludges MessageKludges
}

type MessageKludge struct {
	Name   string
	Value  string
}

type MessageKludges []MessageKludge

func NewNetmailMessage() *NetmailMessage {
	nm := new(NetmailMessage)
	return nm
}

func (self *NetmailMessage) AddKludge(name string, value string) {
	kludge := MessageKludge{
		Name: name,
		Value: value,
	}
	self.Kludges = append(self.Kludges, kludge)
}

func (self *NetmailMessage) SetReply(reply string) {
	self.Reply = reply
}

func (self *NetmailMessage) GetBody() string {
	return self.body
}

func (self *NetmailMessage) SetBody(body string) {
	self.body = body
}
