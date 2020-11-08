package mapper

import uuid "github.com/satori/go.uuid"

const (
	DraftStateActive = 0
	DraftStateDone = 1
)

type Draft struct {
	uuid    string
	subject string
	body    string
	done    int
	orig    string
	id      string
	to      string
	to_addr string
	area    string
	reply   string
	state   int
}

func NewDraft() *Draft {
	newDraft := new(Draft)
	newDraft.uuid = newDraft.makeUUID()
	newDraft.done = DraftStateActive
	return newDraft
}

func (self Draft) makeUUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}

func (self Draft) GetId() string {
	return self.id
}

func (self Draft) GetFrom() string {
	return self.orig
}

func (self *Draft) SetSubject(subj string) {
	self.subject = subj
}

func (self *Draft) SetBody(body string) {
	self.body = body
}

func (self *Draft) SetTo(to string) {
	self.to = to
}

func (self *Draft) SetToAddr(addr string) {
	self.to_addr = addr
}

func (self Draft) GetArea() string {
	return self.area
}

func (self Draft) GetBody() string {
	return self.body
}

func (self Draft) GetTo() string {
	return self.to
}

func (self Draft) GetToAddr() string {
	return self.to_addr
}

func (self Draft) GetSubject() string {
	return self.subject
}

func (self Draft) GetReply() string {
	return self.reply
}

func (self Draft) GetUUID() string {
	return self.uuid
}

func (self *Draft) SetUUID(uuid string) {
	self.uuid = uuid
}

func (self Draft) IsDone() bool {
	return self.state == DraftStateDone
}

func (self *Draft) SetState(state int) {
	self.state = state
}

func (self *Draft) SetArea(area string) {
	self.area = area
}

func (self *Draft) SetId(id string) {
	self.id = id
}

func (self Draft) IsEchoMail() bool {
	return self.area != ""
}

func (self *Draft) SetReply(reply string) {
	self.reply = reply
}

func (self Draft) IsReply() bool {
	return self.reply != ""
}
