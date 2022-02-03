package msg

import (
	"strings"
	"time"
)

type Message struct {
	ID          string
	MsgID       string
	Hash        string
	Area        string
	From        string
	To          string
	Subject     string
	Content     string
	UnixTime    int64
	DateWritten time.Time
	ViewCount   int
	Packet      []byte
	Reply       string
	FromAddr    string
}

func NewMessage() *Message {
	msg := new(Message)
	return msg
}

func (self Message) GetContent() string {
	return self.Content
}

func (self *Message) SetID(id string) {
	self.ID = id
}

func (self *Message) SetMsgID(msgId string) {
	self.MsgID = msgId
}

func (self *Message) SetArea(area string) {
	self.Area = strings.ToUpper(area)
}

func (self *Message) SetFrom(from string) {
	self.From = from
}

func (self *Message) SetTo(to string) {
	self.To = to
}

func (self *Message) SetSubject(subject string) {
	self.Subject = subject
}

func (self *Message) SetContent(content string) {
	self.Content = strings.TrimRight(content, "\x00")
}

func (self *Message) SetUnixTime(unixTime int64) {
	self.UnixTime = unixTime
	self.DateWritten = time.Unix(unixTime, 0)
}

func (self *Message) SetTime(ptm time.Time) {
	self.DateWritten = ptm
	self.UnixTime = ptm.Unix()
}

func (self *Message) SetViewCount(count int) {
	self.ViewCount = count
}

func (self *Message) SetMsgHash(hash string) {
	self.Hash = hash
}

func (self *Message) SetPacket(packet []byte) {
	self.Packet = packet
}

func (self *Message) SetReply(reply string) {
	self.Reply = reply
}

func (self Message) GetFrom() string {
	return self.From
}

func (self Message) GetMsgID() string {
	return self.MsgID
}

func (self *Message) SetFromAddr(addr string) {
	self.FromAddr = addr
}

func (self Message) GetFromAddr() string {
	return self.FromAddr
}

func (self Message) IsNew() bool {
	return self.ViewCount == 0
}
