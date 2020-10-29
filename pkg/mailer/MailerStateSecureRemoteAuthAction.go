package mailer

type MailerStateSecureAuthRemoteAction struct {
	MailerState
}

func NewMailerStateSecureAuthRemoteAction() *MailerStateSecureAuthRemoteAction {
	return new(MailerStateSecureAuthRemoteAction)
}

func (self *MailerStateSecureAuthRemoteAction) String() string {
	return "MailerStateSecureAuthRemoteAction"
}

func (self *MailerStateSecureAuthRemoteAction) Process(mailer *Mailer) IMailerState {

	mailer.stream.WritePassword(mailer.respAuthorization)

	return NewMailerStateIfSecure()

}
