package mailer

import "log"

type ReceiveRoutineResult string

const (
	RxOk       ReceiveRoutineResult = "RxOk"
	RxFailure  ReceiveRoutineResult = "RxFailure"
	RxContinue ReceiveRoutineResult = "RxContinue"
)

func ReceiveRoutine(mailer *Mailer) ReceiveRoutineResult {

	log.Printf("ReceiveRoutine: rxState = %+v", mailer.rxState)

	switch mailer.rxState {
	case RxWaitF:
		return ReceiveRoutineRxWaitF(mailer)

	case RxAccF:
		return ReceiveRoutineRxAccF(mailer)

	case RxRaceD:
		return ReceiveRoutineRxRaceD(mailer)

	case RxWriteD:
		return ReceiveRoutineRxWriteD(mailer)

	case RxEOB:
		return ReceiveRoutineRxEOB(mailer)
	}

	panic("wrong case or memory corruption")

}
