package mailer

func mailerStateOpts(mailer *Mailer) mailerStateFn {
	//mailer.stream.WritePassword("-")
	return mailerStateInitTransfer
}
