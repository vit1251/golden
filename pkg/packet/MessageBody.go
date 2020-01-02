package packet

import (
	"log"
)

type MessageBody struct {
	Kludges   map[string]string
	Body      string
	RAW     []byte
}

func NewMessageBody() (*MessageBody) {
	mb := new(MessageBody)
	mb.Kludges = make(map[string]string, 0)
	return mb
}

func (self *MessageBody) IsEchoMail() bool {
	var result bool = false
	if _, ok := self.Kludges["AREA"]; ok {
		result = true
	}
	return result
}

func (self *MessageBody) SetKludge(name []byte, value []byte) {
	//
	log.Printf("Set kludge: name = %q value = %q", name, value)
	//
	var strName string = string(name)
	var strValue string = string(value)
	//
	self.Kludges[strName] = strValue
}

func (self *MessageBody) GetKludge(name string) string {
	return self.Kludges[name]
}

func (self *MessageBody) SetBody(msg []byte) {
	unicodeBody, _ := DecodeText(msg)
	var body string = string(unicodeBody)
	//
	self.Body = body
}

func (self *MessageBody) Text() string {
	return self.Body
}

