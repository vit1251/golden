package mailer

type MailerStateTxWLA struct {
	MailerState
}

func NewMailerStateTxWLA() *MailerStateTxWLA {
	return new(MailerStateTxWLA)
}

func (self MailerStateTxWLA) String() string {
	return "MailerStateTxWLA"
}

func (self *MailerStateTxWLA) Process(mailer *Mailer) IMailerState {

	if mailer.rxState == RxDone {
		mailer.txState = TxDone
	}

	return NewMailerStateSwitch()

}