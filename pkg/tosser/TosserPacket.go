package tosser

import "github.com/vit1251/golden/pkg/packet"

type TosserPacket struct {
	header  *packet.PacketHeader
	message *packet.PackedMessage
}

func NewTosserPacket() *TosserPacket {
	return new(TosserPacket)
}

func (self *TosserPacket) SetHeader(header  *packet.PacketHeader) {
	self.header = header
}

func (self *TosserPacket) SetMessage(message *packet.PackedMessage) {
	self.message = message
}

func (self *TosserPacket) GetMessage() *packet.PackedMessage {
	return self.message
}

func (self *TosserPacket) GetHeader() *packet.PacketHeader {
	return self.header
}
