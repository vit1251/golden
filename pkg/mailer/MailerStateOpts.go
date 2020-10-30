package mailer

type MailerStateOpts struct {
	MailerState
}

func NewMailerStateOpts() *MailerStateOpts {
	return new(MailerStateOpts)
}

func (self *MailerStateOpts) String() string {
	return "MailerStateOpts"
}

func (self *MailerStateOpts) Process(mailer *Mailer) IMailerState {

	//mailer.stream.WritePassword("-")

	return NewMailerStateInitTransfer()

}
