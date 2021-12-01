package mailer

import (
	"fmt"
	"log"
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

	/* Process queue */
	log.Printf("Process postpone entries in TheQueue")
	for !mailer.queue.IsEmpty() {
		ProcessTheQueue(mailer)
	}

	/* Close session */
	log.Printf("Close stream and session")
	if mailer.stream != nil {
		mailer.stream.CloseSession()
	}

	/* Update status */
	status := fmt.Sprintf("Complete: RX = %+v TX = %+v", mailer.rxRoutineResult, mailer.txRoutineResult)
	mailer.report.SetStatus(status)

	return nil

}
