package msg

import (
	"github.com/xeonx/timeago"
	"strings"
	"time"
)

type MsgAttr uint16

const MsgAttrPrivate       = 1 << 0
const MsgAttrCrash         = 1 << 1
const MsgAttrRecd          = 1 << 2
const MsgAttrSent          = 1 << 3
const MsgAttrFileAttached  = 1 << 4
const MsgAttrInTransit     = 1 << 5
const MsgAttrOrphan        = 1 << 6
const MsgAttrKillSent      = 1 << 7
const MsgAttrLocal         = 1 << 8
const MsgAttrHoldForPickup = 1 << 9
//const MsgAttr       = 1 << 10
const MsgAttrFileRequest   = 1 << 11
const MsgAttrReturnReceiptRequest = 1 << 12
const MsgAttrIsReturnReceipt = 1 << 13
const MsgAttrAuditRequest   = 1 << 14
const FileUpdateReq = 1 << 15

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
	DateWritten *time.Time
	ViewCount   int
	Packet      []byte
}

func NewMessage() *Message {
	msg := new(Message)
	return msg
}

func (self *Message) GetContent() string {
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
	tm := time.Unix(unixTime, 0)
	self.DateWritten = &tm
}

func (self *Message) SetTime(ptm *time.Time) {
	self.DateWritten = ptm
	if ptm != nil {
		var tm time.Time = *ptm
		self.UnixTime = tm.Unix()
	}
}

func (self *Message) SetViewCount(count int) {
	self.ViewCount = count
}

func (self *Message) SetMsgHash(hash string) {
	self.Hash = hash
}

func (self *Message) Age() string {
	var result string = "-"
	if self.DateWritten != nil {
		result = timeago.English.Format(*self.DateWritten)
	}
	return result
}

func (self *Message) SetPacket(packet []byte) {
	self.Packet = packet
}