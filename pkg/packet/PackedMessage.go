package packet

import "github.com/vit1251/golden/pkg/fidotime"

type PackedMessage struct {
	OrigAddr        NetAddr
	DestAddr        NetAddr
	Attributes      uint16
	Time           *fidotime.FidoDate
	ToUserName    []byte
	FromUserName  []byte
	Subject       []byte
	Text          []byte
}

func NewPackedMessage() *PackedMessage {
	msgHeader := new(PackedMessage)
	return msgHeader
}

func (self *PackedMessage) SetAttribute(attr MsgAttr) {
	self.Attributes = uint16(attr)
}

func (self *PackedMessage) SetToUserName(ToUserName []byte) error {
	self.ToUserName = ToUserName
	return nil
}

func (self *PackedMessage) SetFromUserName(FromUserName []byte) error {
	self.FromUserName = FromUserName
	return nil
}

func (self *PackedMessage) SetSubject(s []byte) error {
	self.Subject = s
	return nil
}

func (self *PackedMessage) SetTime(t *fidotime.FidoDate) error {
	self.Time = t
	return nil
}
