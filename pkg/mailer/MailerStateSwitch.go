package mailer

import (
	"log"
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

func (self *MailerStateSwitch) processBoth(mailer *Mailer) IMailerState {

	select {

	/* Data available in Input Buffer */
	case _, ok := <- mailer.stream.InFrameReady:
		log.Printf("Data available in Input Buffer")
		if ok {
			mailer.rxRoutineResult = ReceiveRoutine(mailer)
			return NewMailerStateReceive()
		} else {
			return NewMailerStateEnd()
		}

	/* Free space exists in output buffer */
	case mailer.stream.OutFrameReady <- nil:
		log.Printf("Free space exists in output buffer")
		mailer.txRoutineResult = TransmitRoutine(mailer)
		return NewMailerStateTransmit()

	}

}

func (self *MailerStateSwitch) processReading(mailer *Mailer) IMailerState {

	select {

	/* Data available in Input Buffer */
	case _, ok := <-mailer.stream.InFrameReady:
		log.Printf("Data available in Input Buffer")
		if ok {
			mailer.rxRoutineResult = ReceiveRoutine(mailer)
			return NewMailerStateReceive()
		} else {
			return NewMailerStateEnd()
		}
	}

}

func (self *MailerStateSwitch) Process(mailer *Mailer) IMailerState {

	/* RxState is RxDone and TxState is TxDone */
	if mailer.rxState == RxDone && mailer.txState == TxDone {
		return NewMailerStateEnd()
	}

	/* TxState */
	if mailer.txState == TxWLA {
		return self.processReading(mailer)
	} else {
		return self.processBoth(mailer)
	}

	return self

}
