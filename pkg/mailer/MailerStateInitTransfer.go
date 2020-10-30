package mailer

type MailerStateInitTransfer struct {
	MailerState
}

func NewMailerStateInitTransfer() *MailerStateInitTransfer {
	mscc := new(MailerStateInitTransfer)
	return mscc
}

func (self *MailerStateInitTransfer) String() string {
	return "MailerStateInitTransfer"
}

func (self *MailerStateInitTransfer) Process(mailer *Mailer) IMailerState {

	mailer.rxState = RxWaitF
	mailer.txState = TxGNF

	return NewMailerStateSwitch()

}

