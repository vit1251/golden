package packet

import (
	"github.com/vit1251/golden/pkg/fidotime"
	"log"
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
	newPasswordSize := 8
	newPassword := make([]byte, newPasswordSize)
	copy(newPassword, password)
	log.Printf("newPass: str = %s raw = %+v", newPassword, newPassword)
	self.PktPassword = newPassword
}

type PacketAttr int8

const (
	PacketAttrDirect PacketAttr = 0x01
)

func (self *PacketMessageHeader) UnsetAttribute(attr PacketAttr) error {
	return nil
}

func (self *PacketMessageHeader) SetAttribute(attr PacketAttr) error {
	return nil
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
