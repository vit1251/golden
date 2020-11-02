package mailer

type MailerStateReceive struct {
	MailerState
}

func NewMailerStateReceive() *MailerStateReceive {
	return new(MailerStateReceive)
}

func (self *MailerStateReceive) String() string {
	return "MailerStateReceive"
}

func (self *MailerStateReceive) Process(mailer *Mailer) IMailerState {

	/* Receive routine returned OK */
	if mailer.rxRoutineResult == RxOk {
		return NewMailerStateSwitch()
	}

	/* Receive routine returned Failure */
	if mailer.rxRoutineResult == RxFailure {
		return NewMailerStateEnd()
	}

	if mailer.rxRoutineResult == RxContinue {
		mailer.rxRoutineResult = ReceiveRoutine(mailer)
		return self
	}

	panic("unknown state")

}
