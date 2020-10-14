package mailer

type MailerStateAuthorization struct {
	MailerState
	IMailerState
}

func NewMailerStateAuthorization() *MailerStateAuthorization {
	msa := new(MailerStateAuthorization)
	return msa
}

func (self *MailerStateAuthorization) String() string {
	return "MailerStateAuthorization"
}

func (self *MailerStateAuthorization) Process(mailer *Mailer) IMailerState {

	mailer.WritePassword(mailer.respAuthorization)

	return NewMailerStateSecure()

}