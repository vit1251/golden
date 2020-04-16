package netmail

type NetmailMessage struct {
	ID        string
	MsgID     string
	Hash      string
	From      string
	To        string
	Subject   string
	Content   string
	UnixTime  int64
	ViewCount int
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

func (self *NetmailMessage) SetUnixTime(date int64) {
	self.UnixTime = date
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

