package mailer

func TransmitRoutineTxTryR(mailer *Mailer) TransmitRoutineResult {

	/* Check The Queue */
	if mailer.queue.IsEmpty() {

		/* TheQueue is empty */
		mailer.txState = TxReadS
		return TxContinue

	} else {

		/* TheQueue is not empty */
		ProcessTheQueue(mailer)
		return TxContinue

	}

	panic("unknown case or memory corruption")

}
