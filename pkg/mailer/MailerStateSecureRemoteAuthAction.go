package mailer

func mailerStateSecureAuthRemoteAction(mailer *Mailer) mailerStateFn {
	mailer.stream.WritePassword(mailer.respAuthorization)
	return mailerStateIfSecure
}
