package mailer

type TransmitRoutineResult string

const (
    TxOk       TransmitRoutineResult = "TxOk"
    TxFailure  TransmitRoutineResult = "TxFailure"
    TxContinue TransmitRoutineResult = "TxContinue"
)

func TransmitRoutine(mailer *Mailer) TransmitRoutineResult {
    if mailer.txState == TxGNF {
	return TransmitRoutineTxGNF(mailer)
    } else if mailer.txState == TxWLA {
	return TransmitRoutineTxWLA(mailer)
    } else if mailer.txState == TxTryR {
	return TransmitRoutineTxTryR(mailer)
    } else if mailer.txState == TxReadS {
	return TransmitRoutineTxReadS(mailer)
    } else if mailer.txState == TxDone {
        return TxOk
    } else {
    	panic("wrong case or memory corruption")
    }
}
