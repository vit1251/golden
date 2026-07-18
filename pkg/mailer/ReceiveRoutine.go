package mailer

type ReceiveRoutineResult string

const (
    RxOk       ReceiveRoutineResult = "RxOk"
    RxFailure  ReceiveRoutineResult = "RxFailure"
    RxContinue ReceiveRoutineResult = "RxContinue"
)

func ReceiveRoutine(mailer *Mailer) ReceiveRoutineResult {

    if mailer.rxState == RxWaitF {
	return ReceiveRoutineRxWaitF(mailer)
    } else if mailer.rxState == RxAccF {
	return ReceiveRoutineRxAccF(mailer)
    } else if mailer.rxState == RxRaceD {
	return ReceiveRoutineRxRaceD(mailer)
    } else if mailer.rxState == RxWriteD {
	return ReceiveRoutineRxWriteD(mailer)
    } else if mailer.rxState == RxEOB {
	return ReceiveRoutineRxEOB(mailer)
    } else if mailer.rxState == RxDone {
        return RxOk
    }

    panic("wrong case or memory corruption")

}
