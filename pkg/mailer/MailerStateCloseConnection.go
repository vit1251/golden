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

	/* Close connection */
	if mailer.conn != nil {
		mailer.conn.Close()
		mailer.conn = nil
	}

	return nil
}
