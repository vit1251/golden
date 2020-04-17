package netmail

import (
	"github.com/xeonx/timeago"
	"time"
)

type NetmailMessage struct {
	ID           string
	MsgID        string
	Hash         string
	From         string
	To           string
	Subject      string
	Content      string
	UnixTime     int64
	ViewCount    int
	DateWritten *time.Time
}

func NewNetmailMessage() *NetmailMessage {
	nm := new(NetmailMessage)
	return nm
}

func (self *NetmailMessage) SetMsgID(msgID string) {
	self.MsgID = msgID
}

func (self *NetmailMessage) SetSubject(subject string) {
	self.Subject = subject
}

func (self *NetmailMessage) SetID(id string) {
	self.ID = id
}

func (self *NetmailMessage) SetFrom(from string) {
	self.From = from
}

func (self *NetmailMessage) SetTo(to string) {
	self.To = to
}

func (self *NetmailMessage) SetViewCount(count int) *NetmailMessage {
	self.ViewCount = count
	return self
}

func (self *NetmailMessage) SetContent(body string) *NetmailMessage {
	self.Content = body
	return self
}

func (self *NetmailMessage) SetHash(hash string) *NetmailMessage {
	self.Hash = hash
	return self
}

func (self *NetmailMessage) GetContent() string {
	return self.Content
}

func (self *NetmailMessage) SetUnixTime(unixTime int64) {
	self.UnixTime = unixTime
	tm := time.Unix(unixTime, 0)
	self.DateWritten = &tm
}

func (self *NetmailMessage) SetTime(ptm *time.Time) {
	self.DateWritten = ptm
	if ptm != nil {
		var tm time.Time = *ptm
		self.UnixTime = tm.Unix()
	}
}

func (self *NetmailMessage) Age() string {
	var result string = "-"
	if self.DateWritten != nil {
		result = timeago.English.Format(*self.DateWritten)
	}
	return result
}

