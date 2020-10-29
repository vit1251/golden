package mailer

type MailerStateAuthRemoteAction struct {
	MailerState
}

func NewMailerStateAuthRemoteAction() *MailerStateAuthRemoteAction {
	return new(MailerStateAuthRemoteAction)
}

func (self *MailerStateAuthRemoteAction) String() string {
	return "MailerStateAuthRemoteAction"
}

func (self *MailerStateAuthRemoteAction) Process(mailer *Mailer) IMailerState {

	mailer.stream.WritePassword("-")

	return NewMailerStateIfSecure()

}
