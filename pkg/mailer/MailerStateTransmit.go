package mailer

func mailerStateTransmit(mailer *Mailer) mailerStateFn {

	/* Transmit routine returned OK */
	if mailer.txRoutineResult == TxOk {
		return mailerStateSwitch
	}

	/* Transmit routine returned Failure */
	if mailer.txRoutineResult == TxFailure {
		return mailerStateEnd
	}

	/* Transmit routine returned Continue */
	if mailer.txRoutineResult == TxContinue {
		mailer.txRoutineResult = TransmitRoutine(mailer)
		return mailerStateTransmit
	}

	panic("unkown state")

}
