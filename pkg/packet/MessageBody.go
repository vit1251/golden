package packet

import (
	//	"log"
	"bytes"
	"fmt"
)

type MessageBody struct {
	area    string
	kludges []Kludge
	lines   [][]byte
	origin  []byte
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

func (self MessageBody) GetContent() []byte {
	return bytes.Join(self.lines, []byte(CR))
}

func (self *MessageBody) SetContent(content []byte) {
	rows := bytes.Split(content, []byte(CR))
	for _, row := range rows {
		self.AddLine(row)
	}
}

func (self *MessageBody) SetOrigin(origin []byte) {
	self.origin = origin
}

func (self MessageBody) GetOrigin() []byte {
	return self.origin
}

func (self MessageBody) Bytes() []byte {

	mem := new(bytes.Buffer)

	/* Write AREA section */
	var areaName string = self.GetArea()
	if areaName != "" {
		mem.WriteString(fmt.Sprintf("AREA:%s", areaName))
		mem.WriteString(CR)
	}

	/* Write kludges */
	for _, k := range self.kludges {
		mem.WriteString(fmt.Sprintf("%s", k.Raw))
		mem.WriteString(CR)
	}

	/* Write message body */

	msgBodyRaw := self.GetContent()
	mem.Write(msgBodyRaw)

	return mem.Bytes()
}
