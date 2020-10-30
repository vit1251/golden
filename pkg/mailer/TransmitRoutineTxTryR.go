package mailer

func TransmitRoutineTxTryR(mailer *Mailer) {

	/* Check The Queue */
	if mailer.queue.IsEmpty() {

		/* TheQueue is empty */
		mailer.txState = TxReadS

	} else {

		/* TheQueue is not empty */
		ProcessTheQueue(mailer)

	}

}
