package packet

import (
//	"log"
)

type MessageBody struct {
	Area    string
	kludges []Kludge
	RAW     []byte
}

func NewMessageBody() *MessageBody {
	mb := new(MessageBody)
	return mb
}

func (self *MessageBody) IsEchoMail() bool {
	var result bool = false
	for _, k := range self.kludges {
		if k.Name == "AREA" {
			result = true
			break
		}
	}
	return result
}

func (self *MessageBody) GetArea() (string) {
	return self.Area
}

func (self *MessageBody) SetArea(area string) {
	self.Area = area
}

func (self *MessageBody) AddKludge(name string, value string) {
	newKludge := Kludge{
		Name: name,
		Value: value,
	}
	self.kludges = append(self.kludges, newKludge)
}

func (self *MessageBody) GetKludges() []Kludge {
	return self.kludges
}

func (self *MessageBody) GetKludge(name string, defaultValue string) string {
	var result string = defaultValue
	for _, k := range self.kludges {
		if k.Name == name {
			result = k.Value
			break
		}
	}
	return result
}

func (self *MessageBody) GetRaw() ([]byte) {
	return self.RAW
}

func (self *MessageBody) SetRaw(rawBody []byte) {
	self.RAW = rawBody
}
