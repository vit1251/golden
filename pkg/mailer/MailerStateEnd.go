package mailer

import (
	"log"
	"time"
)

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

	log.Printf("Exit")

	select {
	case <-time.After(5 * time.Second):
		/* Close session */
		if mailer.stream != nil {
			mailer.stream.CloseSession()
		}
		return nil

	}

}
