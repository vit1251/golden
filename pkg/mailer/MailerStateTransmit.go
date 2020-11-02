package mailer

type MailerStateTransmit struct {
	MailerState
}

func NewMailerStateTransmit() *MailerStateTransmit {
	return new(MailerStateTransmit)
}

func (self *MailerStateTransmit) String() string {
	return "MailerStateTransmit"
}

func (self *MailerStateTransmit) Process(mailer *Mailer) IMailerState {

	/* Transmit routine returned OK */
	if mailer.txRoutineResult == TxOk {
		return NewMailerStateSwitch()
	}

	/* Transmit routine returned Failure */
	if mailer.txRoutineResult == TxFailure {
		return NewMailerStateEnd()
	}

	/* Transmit routine returned Continue */
	if mailer.txRoutineResult == TxContinue {
		mailer.txRoutineResult = TransmitRoutine(mailer)
		return self
	}

	panic("unkown state")

}
