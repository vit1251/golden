package tosser

import "strings"

type EchoMessage struct {
	Subject  string
	To       string
	From     string
	body     string
	AreaName string
	Reply    string
	Kludges  MessageKludges
}

func (m *EchoMessage) SetSubject(subj string) {
	m.Subject = subj
}

func (self *EchoMessage) GetBody() string {
	return self.body
}

func (self *EchoMessage) SetBody(body string) {
	newBody := strings.ReplaceAll(body, "\r\n", "\r")
	self.body = newBody
}

func (self *EchoMessage) SetReply(reply string) {
	self.Reply = reply
}

func NewEchoMessage() *EchoMessage {
	em := new(EchoMessage)
	return em
}
