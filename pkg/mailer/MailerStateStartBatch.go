package mailer

type MailerStateStartBatch struct {
	MailerState
}

func NewMailerStateStartBatch() *MailerStateStartBatch {
	mscc := new(MailerStateStartBatch)
	return mscc
}

func (self *MailerStateStartBatch) String() string {
	return "MailerStateStartBatch"
}

func (self *MailerStateStartBatch) Process(mailer *Mailer) IMailerState {

	// TODO - write message ...

	return NewMailerStateReceive()
}

