package mailer

type MailerState struct {
}

type IMailerState interface {
	Process(mailer *Mailer) IMailerState
	String() string
}
