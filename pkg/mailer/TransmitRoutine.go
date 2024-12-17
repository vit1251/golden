package mailer

import "log"

type TransmitRoutineResult string

const (
	TxOk       TransmitRoutineResult = "TxOk"
	TxFailure  TransmitRoutineResult = "TxFailure"
	TxContinue TransmitRoutineResult = "TxContinue"
)

func TransmitRoutine(mailer *Mailer) TransmitRoutineResult {

	log.Printf("TransmitRoutine: txState = %+v", mailer.txState)

	switch mailer.txState {

	case TxGNF:
		return TransmitRoutineTxGNF(mailer)

	case TxWLA:
		return TransmitRoutineTxWLA(mailer)

	case TxTryR:
		return TransmitRoutineTxTryR(mailer)

	case TxReadS:
		return TransmitRoutineTxReadS(mailer)

	}

	panic("wrong case or memory corruption")

}
