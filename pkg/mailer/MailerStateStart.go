package mailer

type MailerStateStart struct {
	MailerState
}

func NewMailerStateStart() *MailerStateStart {
	state := new(MailerStateStart)
	return state
}

func (self MailerStateStart) String() string {
	return "MailerStateStart"
}

func (self *MailerStateStart) Process(mailer *Mailer) IMailerState {
	return NewMailerStateConnInit()
}
