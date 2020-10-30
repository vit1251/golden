package mailer

type MailerStateAuthRemote struct {
	MailerState
}

func NewMailerStateAuthRemote() *MailerStateAuthRemote {
	return new(MailerStateAuthRemote)
}

func (self *MailerStateAuthRemote) String() string {
	return "MailerStateAuthRemote"
}

func (self *MailerStateAuthRemote) Process(mailer *Mailer) IMailerState {

	mailer.stream.WritePassword("-")

	return NewMailerStateIfSecure()

}
