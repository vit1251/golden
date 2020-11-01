package packet

import (
	"github.com/vit1251/golden/pkg/fidotime"
)

type PktVersion int

const (
	PKTv2      = 0x0200
	PKTv2plus  = 0x0201
	PKTv22     = 0x0202
)

func NewPacketHeader() *PacketHeader {
	ph := new(PacketHeader)
	return ph
}

func (self *PacketHeader) SetPassword(password string) {
	newPassword := make([]byte, 8)
	copy(newPassword, password)
	self.PktPassword = newPassword
}

func (self *PacketMessageHeader) SetAttribute(attr MsgAttr) {
	self.Attributes = uint16(attr)
}

func (self *PacketMessageHeader) SetToUserName(ToUserName []byte) error {
	self.ToUserName = ToUserName
	return nil
}

func (self *PacketMessageHeader) SetFromUserName(FromUserName []byte) error {
	self.FromUserName = FromUserName
	return nil
}

func (self *PacketMessageHeader) SetSubject(s []byte) error {
	self.Subject = s
	return nil
}

func (self *PacketMessageHeader) SetTime(t *fidotime.FidoDate) error {
	self.Time = t
	return nil
}
