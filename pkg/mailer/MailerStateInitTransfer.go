package mailer

func mailerStateInitTransfer(mailer *Mailer) mailerStateFn {

	mailer.rxState = RxWaitF
	mailer.txState = TxGNF

	return mailerStateSwitch

}
