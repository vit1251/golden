package packet

import (
	"bytes"
	"log"
)

type MessageBodyAttach struct {
	permission	string            /* Permission    */
	name		string            /* Name          */
	uue		[]string          /* UUE rows      */
	buffer		bytes.Buffer      /* raw result    */
}

func NewMessageBodyAttach() *MessageBodyAttach {
	return new(MessageBodyAttach)
}

func (self *MessageBodyAttach) Write(buf []byte) error {
	self.buffer.Write(buf)
	return nil
}

func decodeByte(b byte) byte {
	r := byte(b) - 32
	r = r & 077
	return r
}

func extractBlockCount(row []byte) byte {
	rowSize := len(row)
	if rowSize > 0 {
		return decodeByte(row[0])
	}
	return 0
}

func (self *MessageBodyAttach) WriteLine(row []byte) error {

	/* Save UUE source */
	line := string(row)
	self.uue = append(self.uue, line)

	/* Extract UUE section count */
	n := extractBlockCount(row)
	log.Printf("UUE: blockCount = %d", n)

	var rawData []byte
//	var rowSize int = len(row)
	var p int = 1

	for n > 0 {

		rawData = row[p:]
		log.Printf("rawData = %s", rawData)

		if n >= 3 {
			b1 := decodeByte(rawData[0]) << 2 | decodeByte(rawData[1]) >> 4
			self.buffer.WriteByte(b1)
			b2 := decodeByte(rawData[1]) << 4 | decodeByte(rawData[2]) >> 2
			self.buffer.WriteByte(b2)
			b3 := decodeByte(rawData[2]) << 6 | decodeByte(rawData[3])
			self.buffer.WriteByte(b3)
		} else {
			if n >= 1 {
				b1 := decodeByte(rawData[0]) << 2 | decodeByte(rawData[1]) >> 4
				self.buffer.WriteByte(b1)
			}
			if n >= 2 {
				b2 := decodeByte(rawData[1]) << 4 | decodeByte(rawData[2]) >> 2
				self.buffer.WriteByte(b2)
			}
		}

		n = n - 3
		p = p + 4

	}

	return nil

}

func (self *MessageBodyAttach) SetPermission(permission string) {
	self.permission = permission
}

func (self *MessageBodyAttach) SetName(name string) {
	self.name = name
}

func (self *MessageBodyAttach) Len() int {
	return self.buffer.Len()
}

func (self *MessageBodyAttach) GetData() bytes.Buffer {
	return self.buffer
}

func (self *MessageBodyAttach) GetName() string {
	return self.name
}
