package mailer

func mailerStateReceive(mailer *Mailer) mailerStateFn {

	/* Receive routine returned OK */
	if mailer.rxRoutineResult == RxOk {
		return mailerStateSwitch
	}

	/* Receive routine returned Failure */
	if mailer.rxRoutineResult == RxFailure {
		return mailerStateEnd
	}

	if mailer.rxRoutineResult == RxContinue {
		mailer.rxRoutineResult = ReceiveRoutine(mailer)
		return mailerStateReceive
	}

	panic("unknown state")

}
