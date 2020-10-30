package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func ReceiveRoutineRxRaceD(mailer *Mailer) {

	/* Accept from beginning */
	nextFrame := <- mailer.stream.InFrame

	if nextFrame.IsDataFrame() {

		packet := nextFrame.DataFrame.Body

		log.Printf("Data frame: body = %d", len(packet))

		if mailer.recvStream != nil {

			mailer.recvStream.Write(packet)

		} else {
			log.Printf("No write stream - skip data packet")
		}

		mailer.rxState = RxWriteD
	}

	/* Got M_ERR */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_ERR {
			/* Report Error */
			log.Printf("MailerState: Receive - RxRaceD - Got M_ERR")
			mailer.rxState = RxDone
		}
	}

	/* Got M_GET / M_GOT / M_SKIP */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_GET || nextFrame.CommandID == stream.M_GOT || nextFrame.CommandID == stream.M_SKIP {
			/* Add frame to The Queue */
			mailer.queue.Push(nextFrame)
		}
	}

	/* Got M_NUL */
	// TODO -

	/* Got M_FILE */
	// TODO -

	/* Got other known frame */
	// TODO -

	// TODO - Report receiving file

}

