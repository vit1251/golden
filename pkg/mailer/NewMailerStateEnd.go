package mailer

import "log"

type MailerStateEnd struct {
	MailerState
}

func NewMailerStateEnd() *MailerStateEnd {
	state := new(MailerStateEnd)
	return state
}

func (self *MailerStateEnd) String() string {
	return "MailerStateEnd"
}

func (self *MailerStateEnd) Process(mailer *Mailer) IMailerState {
	log.Printf("End")
	return nil
}
