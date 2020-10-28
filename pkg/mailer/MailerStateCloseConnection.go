package mailer

type MailerStateCloseConnection struct {
	MailerState
}

func NewMailerStateCloseConnection() *MailerStateCloseConnection {
	mscc := new(MailerStateCloseConnection)
	return mscc
}

func (self *MailerStateCloseConnection) String() string {
	return "MailerStateCloseConnection"
}

func (self *MailerStateCloseConnection) Process(mailer *Mailer) IMailerState {

	close(mailer.inStop)
	close(mailer.outStop)

	return nil
}
