package packet

import (
//	"log"
	"bytes"
)

type MessageBody struct {
	area    string
	kludges []Kludge
	lines   [][]byte
}

func NewMessageBody() *MessageBody {
	mb := new(MessageBody)
	return mb
}

func (self MessageBody) IsArea() bool {
	return self.area != ""
}

func (self MessageBody) GetArea() string {
	return self.area
}

func (self *MessageBody) SetArea(area string) {
	self.area = area
}

func (self *MessageBody) AddKludge(k Kludge) {
	self.kludges = append(self.kludges, k)
}

func (self *MessageBody) GetKludges() []Kludge {
	return self.kludges
}

func (self *MessageBody) AddLine(line []byte) {
	self.lines = append(self.lines, line)
}

func (self *MessageBody) GetRaw() []byte {
	return bytes.Join(self.lines, []byte(CR))
}

func (self *MessageBody) SetRaw(rawBody []byte) {
	rows := bytes.Split(rawBody, []byte(CR))
	for _, row := range rows {
		self.AddLine(row)
	}
}
