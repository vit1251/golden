package mailer

import (
	"fmt"
	"log"
)

func mailerStateEnd(mailer *Mailer) mailerStateFn {

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
