package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func ReceiveRoutineRxRaceD(mailer *Mailer) ReceiveRoutineResult {

	/* Get a frame from Input Buffer */
	nextFrame := <-mailer.stream.InFrame

	/* Got Data frame */
	if nextFrame.IsDataFrame() {

		packet := nextFrame.DataFrame.Body

		log.Printf("Data frame: body = %d", len(packet))

		if mailer.recvStream != nil {

			mailer.recvStream.Write(packet)

		} else {
			log.Printf("No write stream - skip data packet")
		}

		mailer.rxState = RxWriteD
		return RxContinue
	}

	/* Got M_ERR */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_ERR {
			/* Report Error */
			log.Printf("MailerState: Receive - RxRaceD - Got M_ERR")

			mailer.rxState = RxDone
			return RxFailure
		}
	}

	/* Got M_GET / M_GOT / M_SKIP */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_GET || nextFrame.CommandID == stream.M_GOT || nextFrame.CommandID == stream.M_SKIP {
			/* Add frame to The Queue */
			mailer.queue.Push(nextFrame)

			mailer.rxState = RxRaceD
			return RxOk
		}
	}

	/* Got M_NUL */
	// TODO -

	/* Got M_FILE */
	// TODO -

	/* Got other known frame */
	if nextFrame.IsCommandFrame() {
		mailer.rxState = RxDone
		return RxFailure
	}

	/* Got unknown frame */
	mailer.rxState = RxRaceD
	return RxOk

}
