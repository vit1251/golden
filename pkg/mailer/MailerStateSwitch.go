package mailer

import (
	"log"
	"time"
)

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

	log.Printf("                         --- Debug wait --- ")
	log.Printf("rxState = %s txState = %s", mailer.rxState, mailer.txState)
	time.Sleep(1 * time.Second)

	/* Check complete */
	log.Printf("Check complete")
	if mailer.rxState == RxDone && mailer.txState == TxDone {
		return NewMailerStateEnd()
	}

	select {

	/* Data available in Input Buffer */
	case _, ok := <- mailer.stream.InFrameReady:
		log.Printf("Data available in Input Buffer")
		if ok {
			ReceiveRoutine(mailer)
		} else {
			// Close session ...
			return NewMailerStateEnd()
		}

	/* Data available in Output Buffer */
	case mailer.stream.OutFrameReady <- nil:
		log.Printf("Data available in Output Buffer")
		TransmitRoutine(mailer)

	}

	return self

}
