package mailer

import (
)

type MailerStateIfSecure struct {
	MailerState
}

func NewMailerStateIfSecure() *MailerStateIfSecure {
	mss := new(MailerStateIfSecure)
	return mss
}

func (self *MailerStateIfSecure) String() string {
	return "MailerStateIfSecure"
}

func (self *MailerStateIfSecure) Process(mailer *Mailer) IMailerState {
	return NewMailerStateWaitOk()
}
