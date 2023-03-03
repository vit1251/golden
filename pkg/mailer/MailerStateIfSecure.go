package mailer

func mailerStateIfSecure(mailer *Mailer) mailerStateFn {
	return mailerStateWaitOk
}
