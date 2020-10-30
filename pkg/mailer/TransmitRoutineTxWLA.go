package mailer

func TransmitRoutineTxWLA(mailer *Mailer) {

	if mailer.queue.IsEmpty() {

		/* The Queue is empty and RxState >= RxEOB */
		if mailer.rxState == RxEOB || mailer.rxState == RxDone {

			mailer.txState = TxDone

		} else {

			/* The Queue is empty and RxState < RxEOB */

		}

	} else {

		/* The Queue is not empty */
		ProcessTheQueue(mailer)

	}

}
