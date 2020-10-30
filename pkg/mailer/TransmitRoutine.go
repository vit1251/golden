package mailer

import "log"

func TransmitRoutine(mailer *Mailer) {

	log.Printf("TransmitRoutine: txState = %+v", mailer.txState)

	switch mailer.txState {

	case TxGNF:
		TransmitRoutineTxGNF(mailer)

	case TxWLA:
		TransmitRoutineTxWLA(mailer)

	case TxTryR:
		TransmitRoutineTxTryR(mailer)

	case TxReadS:
		TransmitRoutineTxReadS(mailer)

	}

}
