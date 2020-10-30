package mailer

type MailerStateTransmit struct {
	MailerState
}

func NewMailerStateTransmit() *MailerStateTransmit {
	return new(MailerStateTransmit)
}

func (self MailerStateTransmit) String() string {
	return "MailerStateTransmit"
}

func (self *MailerStateTransmit) Process(mailer *Mailer) IMailerState {

	if mailer.txState == TxGNF {
		return NewMailerStateTxGNF()
	}

	if mailer.txState == TxWLA {
		return NewMailerStateTxWLA()
	}

	if mailer.txState == TxTryR {
		return NewMailerStateTxTryR()
	}

	if mailer.txState == TxReadS {
		return NewMailerStateTxReadS()
	}

	return NewMailerStateSwitch()

}
