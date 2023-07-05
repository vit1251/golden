package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func mailerStateAdditionalStepProcessCommandFrame(mailer *Mailer, nextFrame stream.Frame) mailerStateFn {

	var streamCommandId = nextFrame.CommandFrame.CommandID

	/* Use modern secure authorization */
	if streamCommandId == stream.M_NUL {
		mailer.processNulFrame(nextFrame)
	}

	/* Use unsecure password authorization */
	if streamCommandId == stream.M_ADR {

		log.Printf("Mailer: Remote address is %+v", nextFrame.CommandFrame.Body)
		mailer.report.SetRemoteIdent(string(nextFrame.CommandFrame.Body))

		if mailer.respAuthorization != "" {
			return mailerStateSecureAuthRemoteAction
		} else {
			return mailerStateAuthRemote
		}
	}

	return mailerStateAdditionalStep
}

func mailerStateAdditionalStepProcessFrame(mailer *Mailer, nextFrame stream.Frame) mailerStateFn {

	if nextFrame.IsCommandFrame() {
		return mailerStateAdditionalStepProcessCommandFrame(mailer, nextFrame)
	}

	return mailerStateAdditionalStep

}

func mailerStateAdditionalStep(mailer *Mailer) mailerStateFn {

	select {
	case nextFrame := <-mailer.stream.InFrame:
		return mailerStateAdditionalStepProcessFrame(mailer, nextFrame)
	}

}
