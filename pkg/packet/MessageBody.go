package packet

import (
//	"log"
)

type Kludge struct {
	Name string
	Value string
}

type MessageBody struct {
	Area      string
	Kludges []Kludge
	RAW     []byte
}

func NewMessageBody() *MessageBody {
	mb := new(MessageBody)
	return mb
}

func (self *MessageBody) IsEchoMail() bool {
	var result bool = false
	for _, k := range self.Kludges {
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
	self.Kludges = append(self.Kludges, newKludge)
}

func (self *MessageBody) GetKludge(name string, defaultValue string) string {
	var result string = defaultValue
	for _, k := range self.Kludges {
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
