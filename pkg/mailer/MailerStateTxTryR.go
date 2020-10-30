package mailer

type MailerStateTxTryR struct {
	MailerState
}

func NewMailerStateTxTryR() *MailerStateTxTryR {
	return new(MailerStateTxTryR)
}

func (self MailerStateTxTryR) String() string {
	return "MailerStateTxTryR"
}

func (self *MailerStateTxTryR) Process(mailer *Mailer) IMailerState {

	/* Check The Queue */
	// TODO - ...

	/* TheQueue is empty */
	mailer.txState = TxReadS

	/* TheQueue is not empty */
	// TODO - call ProcessTheQueue

	return NewMailerStateSwitch()

}
