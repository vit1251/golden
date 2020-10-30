package mailer

type MailerStateSwitch struct {

}

func NewMailerStateSwitch() *MailerStateSwitch {
	msr := new(MailerStateSwitch)
	return msr
}

func (self *MailerStateSwitch) String() string {
	return "MailerStateSwitch"
}

func (self *MailerStateSwitch) Process(mailer *Mailer) IMailerState {


	/* Check complete */
	if mailer.rxState == RxDone && mailer.txState == TxDone {
		return NewMailerStateEnd()
	}

	/* Data available in Input Buffer */
	if mailer.rxState != RxDone {
		return NewMailerStateReceive()
	}

	/* Data available in Output Buffer */
	if mailer.txState != TxDone {
		return NewMailerStateTransmit()
	}

	return self

}


