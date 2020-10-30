package mailer

func TransmitRoutineTxWLA(mailer *Mailer) TransmitRoutineResult {

	/* Check TheQueue */
	if mailer.queue.IsEmpty() {

		/* The Queue is empty and RxState >= RxEOB */
		if mailer.rxState == RxEOB || mailer.rxState == RxDone {

			mailer.txState = TxDone
			return TxOk

		} else {

			/* The Queue is empty and RxState < RxEOB */
			mailer.txState = TxWLA
			return TxOk

		}

	} else {

		/* The Queue is not empty */
		ProcessTheQueue(mailer)
		return TxContinue

	}

	panic("unknown case or memory ocrruption")

}
