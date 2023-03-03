package mailer

func mailerStateAuthRemote(mailer *Mailer) mailerStateFn {
	mailer.stream.WritePassword("-")
	return mailerStateIfSecure
}
