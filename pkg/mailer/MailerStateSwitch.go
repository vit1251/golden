package mailer

import (
	"log"
)

func mailerStateSwitchProcessBoth(mailer *Mailer) mailerStateFn {

	select {

	/* Data available in Input Buffer */
	case _, ok := <-mailer.stream.InFrameReady:
		log.Printf("Data available in Input Buffer")
		if ok {
			mailer.rxRoutineResult = ReceiveRoutine(mailer)
			return mailerStateReceive
		} else {
			return mailerStateEnd
		}

	/* Free space exists in output buffer */
	case mailer.stream.OutFrameReady <- nil:
		log.Printf("Free space exists in output buffer")
		mailer.txRoutineResult = TransmitRoutine(mailer)
		return mailerStateTransmit

	}

}

func mailerStateSwitchProcessReading(mailer *Mailer) mailerStateFn {

	select {

	/* Data available in Input Buffer */
	case _, ok := <-mailer.stream.InFrameReady:
		log.Printf("Data available in Input Buffer")
		if ok {
			mailer.rxRoutineResult = ReceiveRoutine(mailer)
			return mailerStateReceive
		} else {
			return mailerStateEnd
		}
	}

}

func mailerStateSwitch(mailer *Mailer) mailerStateFn {

	/* RxState is RxDone and TxState is TxDone */
	if mailer.rxState == RxDone && mailer.txState == TxDone {
		mailer.report.SetStatus("Session RX/TX complete.")
		return mailerStateEnd
	}

	/* TxState */
	if mailer.txState == TxWLA {
		return mailerStateSwitchProcessReading(mailer)
	} else {
		return mailerStateSwitchProcessBoth(mailer)
	}

	return mailerStateSwitch

}
