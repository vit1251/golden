package msg

import (
	"bytes"
	"github.com/vit1251/golden/pkg/packet"
)

type MessageContent struct {
	rows    [][]byte
	kludges []packet.Kludge
	area    string
	raw     []byte
	packet  []byte
	origin  []byte
}

func NewMessageContent() *MessageContent {
	mc := new(MessageContent)
	return mc
}

func (self *MessageContent) AddLine(line []byte) {
	if bytes.HasPrefix(line, []byte(packet.SOH)) {
		k := packet.NewKludge()
		k.Set(line)
		self.kludges = append(self.kludges, *k)
	}
	self.rows = append(self.rows, line)
}

func (self MessageContent) GetKludges() []packet.Kludge {
	return self.kludges
}

func (self *MessageContent) AddKludge(k packet.Kludge) {
	self.kludges = append(self.kludges, k)
}

func (self MessageContent) GetContent() []byte {
	return bytes.Join(self.rows, []byte(CR))
}

func (self *MessageContent) SetArea(name string) {
	self.area = name
}

func (self MessageContent) IsArea() bool {
	return self.area != ""
}

func (self MessageContent) GetArea() string {
	return self.area
}

func (self *MessageContent) SetPacket(content []byte) {
	self.packet = content
}

func (self MessageContent) GetPacket() []byte {
	return self.packet
}

func (self *MessageContent) SetOrigin(origin []byte) {
	self.origin = origin
}

func (self MessageContent) GetOrigin() []byte {
	return self.origin
}
