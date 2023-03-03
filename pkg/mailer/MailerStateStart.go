package mailer

func mailerStateStart(mailer *Mailer) mailerStateFn {
    return mailerStateConnInit
}
