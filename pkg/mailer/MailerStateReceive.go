package mailer

type MailerStateReceive struct {
	MailerState
}

func NewMailerStateReceive() *MailerStateReceive {
	return new(MailerStateReceive)
}

func (self MailerStateReceive) String() string {
	return "MailerStateReceive"
}

func (self *MailerStateReceive) Process(mailer *Mailer) IMailerState {

	if mailer.rxState == RxWaitF {
		return NewMailerStateRxWaitF()
	}

	if mailer.rxState == RxAccF {
	  	return NewMailerStateRxAccF()
	}

	if mailer.rxState == RxRaceD {
		return NewMailerStateRxRaceD()
	}

	if mailer.rxState == RxWriteD {
		return NewMailerStateRxWriteD()
	}

	if mailer.rxState == RxEOB {
		return NewMailerStateRxEOB()
	}

	return NewMailerStateSwitch()

}
