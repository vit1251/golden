package mapper

import (
	commonfunc "github.com/vit1251/golden/pkg/common"
	"time"
)

type NetmailMsg struct {
	ID          string
	MsgID       string
	Hash        string
	From        string
	To          string
	OrigAddr    string
	DestAddr    string
	Subject     string
	Content     string
	UnixTime    int64
	ViewCount   int
	DateWritten time.Time
	Packet      []byte
}

func NewNetmailMsg() *NetmailMsg {
	nm := new(NetmailMsg)
	return nm
}

func (self *NetmailMsg) SetMsgID(msgID string) {
	self.MsgID = msgID
}

func (self *NetmailMsg) SetSubject(subject string) {
	self.Subject = subject
}

func (self *NetmailMsg) SetID(id string) {
	self.ID = id
}

func (self *NetmailMsg) SetFrom(from string) {
	self.From = from
}

func (self *NetmailMsg) SetTo(to string) {
	self.To = to
}

func (self *NetmailMsg) SetViewCount(count int) *NetmailMsg {
	self.ViewCount = count
	return self
}

func (self *NetmailMsg) SetContent(body string) *NetmailMsg {
	self.Content = body
	return self
}

func (self *NetmailMsg) SetHash(hash string) *NetmailMsg {
	self.Hash = hash
	return self
}

func (self *NetmailMsg) GetContent() string {
	return self.Content
}

func (self *NetmailMsg) SetUnixTime(unixTime int64) {
	self.UnixTime = unixTime
	self.DateWritten = time.Unix(unixTime, 0)
}

func (self *NetmailMsg) SetTime(ptm time.Time) {
	self.DateWritten = ptm
	self.UnixTime = ptm.Unix()
}

func (self NetmailMsg) GetAge() string {
	result := commonfunc.MakeHumanTime(self.DateWritten)
	return result
}

func (self *NetmailMsg) SetPacket(packet []byte) {
	self.Packet = packet
}

func (self *NetmailMsg) SetOrigAddr(addr string) {
	self.OrigAddr = addr
}

func (self *NetmailMsg) SetDestAddr(addr string) {
	self.DestAddr = addr
}

