package mailer

import (
	"log"
	"os"
)

type MailerStateRxAccF struct {
	MailerState
}

func NewMailerStateRxAccF() *MailerStateRxAccF {
	return new(MailerStateRxAccF)
}

func (self MailerStateRxAccF) String() string {
	return "MailerStateRxAccF"
}

func (self *MailerStateRxAccF) Process(mailer *Mailer) IMailerState {

	/* Accept from beginning */

	log.Printf("Open path: %+v", mailer.recvName)

	if stream, err1 := os.Create(mailer.recvName.AbsolutePath); err1 != nil {
		// TODO - erro report and skip ...
	} else {
		mailer.recvStream = stream
	}

	// TODO - Report receiving file

	mailer.rxState = RxRaceD

	return NewMailerStateSwitch()

}
