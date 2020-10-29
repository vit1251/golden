package mailer

import "github.com/vit1251/golden/pkg/mailer/stream"

type MailerStateEndBatch struct {
	MailerState
}

func NewMailerStateEndBatch() *MailerStateEndBatch {
	mscc := new(MailerStateEndBatch)
	return mscc
}

func (self *MailerStateEndBatch) String() string {
	return "MailerStateEndBatch"
}

func (self *MailerStateEndBatch) Process(mailer *Mailer) IMailerState {

	mailer.stream.WriteCommandPacket(stream.M_EOB, []byte("Complete!"))

	return NewMailerStateCloseConnection()
}
