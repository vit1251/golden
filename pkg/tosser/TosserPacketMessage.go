package tosser

import "github.com/vit1251/golden/pkg/packet"

type TosserPacketMessage struct {
//	Packet    *TosserPacket
	Header    *packet.PacketMessageHeader
	Body     []byte
}

func NewTosserPacketMessage() *TosserPacketMessage {
	return new(TosserPacketMessage)
}
