package mailer

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

	mailer.writeCommandPacket(M_EOB, []byte("Complete!"))

	return NewMailerStateCloseConnection()
}
