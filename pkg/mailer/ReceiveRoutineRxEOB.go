package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func ReceiveRoutineRxEOB(mailer *Mailer) ReceiveRoutineResult {

	/* Get a frame from Input Buffer */
	nextFrame := <- mailer.stream.InFrame

	/* Pending Files list is empty */
	if mailer.pendingFiles.IsEmpty() {
		mailer.rxState = RxDone
		return RxOk
	}

	/* Didn't get a complete frame yet or TxState is not TxDone */
	if mailer.txState != TxDone {
		mailer.rxState = RxEOB
		return RxOk
	}

	/* Got M_ERR */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_ERR {
			/* Report Rrror */
			log.Printf("Receive - RxEOB - Got M_ERR")
			mailer.rxState = RxDone
			return RxFailure
		}
	}

	/* Got M_GET / M_GOT / M_SKIP */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_GET || nextFrame.CommandID == stream.M_GOT || nextFrame.CommandID == stream.M_SKIP {
			mailer.queue.Push(nextFrame)
			mailer.rxState = RxEOB
			return RxOk
		}
	}

	/* Got M_NUL */
	// TODO -

	/* Got other known frame or data frame */
	if nextFrame.IsCommandFrame() {
		mailer.rxState = RxDone
		return RxFailure
	}

	/* Got unknown frame */
	mailer.rxState = RxEOB
	return RxOk

}
