package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type MailerStateRxEOB struct {
	MailerState
}

func NewMailerStateRxEOB() *MailerStateRxEOB {
	return new(MailerStateRxEOB)
}

func (self MailerStateRxEOB) String() string {
	return "MailerStateRxEOB"
}

func (self *MailerStateRxEOB) Process(mailer *Mailer) IMailerState {

	/* Get a frame from Input Buffer */
	nextFrame := mailer.stream.GetFrame()

	/* Pending Files list is empty */
	mailer.rxState = RxDone

	/* Didn't get a complete frame yet or TxState is not TxDone */
	// TODO -

	/* Got M_ERR */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_ERR {
			/* Report Rrror */
			log.Printf("Receive - RxEOB - Got M_ERR")
			mailer.rxState = RxDone
		}
	}

	/* Got M_GET / M_GOT / M_SKIP */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_GET || nextFrame.CommandID == stream.M_GOT || nextFrame.CommandID == stream.M_SKIP {
			// TODO - Add frame to The Queue
		}
	}

	/* Got M_NUL */
	// TODO -

	/* Got other known frame or data frame */
	// TODO -

	/* Got unknown frame */
	// TODO -

	return NewMailerStateSwitch()

}