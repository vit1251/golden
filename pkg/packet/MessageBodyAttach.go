package packet

import "bytes"

type MessageBodyAttach struct {
	name string
	permission string
	buffer bytes.Buffer
}

func NewMessageBodyAttach() *MessageBodyAttach {
	return new(MessageBodyAttach)
}

func (self *MessageBodyAttach) Write(buf []byte) error {
	self.buffer.Write(buf)
	return nil
}

func (self *MessageBodyAttach) WriteLine(row []byte) {
	self.Write(row[1:])
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
	var newOut bytes.Buffer
	newSize := self.buffer.Len()
	for i := 0; i < newSize; i += 4 {
		//
		in0, _ := self.buffer.ReadByte() // [i+0]
		in1, _ := self.buffer.ReadByte() // [i+1]
		in2, _ := self.buffer.ReadByte() // [i+2]
		in3, _ := self.buffer.ReadByte() // [i+3]
		//
		in0 = in0 - 32
		in1 = in1 - 32
		in2 = in2 - 32
		in3 = in3 - 32
		//
		out0 := (in0 << 2) | ((0x30 & in1) >> 4)
		out1 := (in1 << 4) | ((0x3c & in2) >> 2)
		out2 := (in2 << 6) | (0x3f & in3)
		//
		newOut.Write([]byte{out0, out1, out2})
		//
	}
	return newOut
}
