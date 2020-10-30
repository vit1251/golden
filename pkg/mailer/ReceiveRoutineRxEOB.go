package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func ReceiveRoutineRxEOB(mailer *Mailer) {

	/* Get a frame from Input Buffer */
	nextFrame := <- mailer.stream.InFrame

	/* Pending Files list is empty */
	if mailer.pendingFiles.IsEmpty() {
		mailer.rxState = RxDone
	}

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
			mailer.queue.Push(nextFrame)
		}
	}

	/* Got M_NUL */
	// TODO -

	/* Got other known frame or data frame */
	// TODO -

	/* Got unknown frame */
	// TODO -

}
