package mailer

import (
	stream2 "github.com/vit1251/golden/pkg/mailer/stream"
	"log"
	"os"
)

const (
	AcceptFromBegin  = 1
	AcceptFromOffset = 2
	AcceptLater      = 3
)

func ReceiveRoutineRxAccF(mailer *Mailer) ReceiveRoutineResult {

	acceptMode := AcceptFromBegin

	/* Accept from beginning */
	if acceptMode == AcceptFromBegin {

		/* Report receiving file */
		log.Printf("Report receiving file: name = %s", mailer.recvName.AbsolutePath)

		stream, err1 := os.Create(mailer.recvName.AbsolutePath)

		if err1 != nil {
			log.Printf("Fail to create file")
			// TODO - error report and packet ...
			mailer.stream.WriteCommandPacket(stream2.M_ERR, []byte("Unable to open file!"))
			mailer.rxState = RxDone
			return RxFailure
		}

		mailer.recvStream = stream
		mailer.rxState = RxRaceD

		return RxOk

	}

	return RxFailure

}
