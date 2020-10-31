package packet

import "github.com/vit1251/golden/pkg/fidotime"

type PacketMessageHeader struct {
	OrigAddr        NetAddr
	DestAddr        NetAddr
	Attributes      uint16
	ToUserName    []byte
	FromUserName  []byte
	Subject       []byte
	Time           *fidotime.FidoDate
}

func NewPacketMessageHeader() *PacketMessageHeader {
	msgHeader := new(PacketMessageHeader)
	return msgHeader
}

