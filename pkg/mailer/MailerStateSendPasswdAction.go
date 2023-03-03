package mailer

func mailerSendPasswdAction(mailer *Mailer) mailerStateFn {
	mailer.stream.WritePassword("-")
	return mailerStateIfSecure
}
