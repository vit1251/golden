package mailer

type MailerSendPasswdAction struct {
	MailerState
}

func NewMailerStateSendPasswdAction() *MailerSendPasswdAction {
	return new(MailerSendPasswdAction)
}

func (self *MailerSendPasswdAction) String() string {
	return "MailerSendPasswdAction"
}

func (self *MailerSendPasswdAction) Process(mailer *Mailer) IMailerState {
	mailer.stream.WritePassword("-")
	return NewMailerStateIfSecure()
}
