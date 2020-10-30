package mailer

import "log"

func ReceiveRoutine(mailer *Mailer) {

	log.Printf("ReceiveRoutine: rxState = %+v", mailer.rxState)

	switch mailer.rxState {
	case RxWaitF:
		ReceiveRoutineRxWaitF(mailer)

	case RxAccF:
		ReceiveRoutineRxAccF(mailer)

	case RxRaceD:
		ReceiveRoutineRxRaceD(mailer)

	case RxWriteD:
		ReceiveRoutineRxWriteD(mailer)

	case RxEOB:
		ReceiveRoutineRxEOB(mailer)
	}

}
