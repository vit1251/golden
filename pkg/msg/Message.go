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

func (m *Message) GetContent() string { return m.Content }
func (m *Message) SetID(id string) { m.ID = id }
func (m *Message) SetMsgID(msgId string) { m.MsgID = msgId }
func (m *Message) SetArea(area string) { m.Area = strings.ToUpper(area) }
func (m *Message) SetFrom(from string) { m.From = from }
func (m *Message) SetTo(to string) { m.To = to }
func (m *Message) SetSubject(subject string) { m.Subject = subject }
func (m *Message) SetContent(content string) { m.Content = content }

func (m *Message) SetUnixTime(unixTime int64) {
    m.UnixTime = unixTime
    m.DateWritten = time.Unix(unixTime, 0)
}

func (m *Message) SetTime(ptm time.Time) {
    m.DateWritten = ptm
    m.UnixTime = ptm.Unix()
}

func (m *Message) SetViewCount(count int) { m.ViewCount = count }
func (m *Message) SetMsgHash(hash string) { m.Hash = hash }
func (m *Message) SetPacket(packet []byte) { m.Packet = packet }
func (m *Message) SetReply(reply string) { m.Reply = reply }
func (m *Message) GetFrom() string { return m.From }
func (m *Message) GetMsgID() string { return m.MsgID }
func (m *Message) SetFromAddr(addr string) { m.FromAddr = addr }
func (m *Message) GetFromAddr() string { return m.FromAddr }
func (m *Message) IsNew() bool { return m.ViewCount == 0 }
func (m *Message) GetPacket() []byte { return m.Packet }
