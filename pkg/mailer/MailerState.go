package mailer

type mailerStateFn func(mailer *Mailer) mailerStateFn
