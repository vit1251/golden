package mailer

import (
	"bytes"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type AdditionalStepAction struct {
	MailerState
}

func NewMailerStateAdditionalStepAction() *AdditionalStepAction {
	return new(AdditionalStepAction)
}

func (self *AdditionalStepAction) String() string {
	return "AdditionalStepAction"
}

func (self *AdditionalStepAction) processCommandFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {

	var streamCommandId = nextFrame.CommandFrame.CommandID

	/* Use modern secure authorization */
	if streamCommandId == stream.M_NUL {
		if key, value, err1 := mailer.parseInfoFrame(nextFrame.CommandFrame.Body); err1 == nil {
			log.Printf("Remote side option: name = %s value = %s", key, value)
			if bytes.Equal(key, []byte("OPT")) {
				mailer.parseInfoOptFrame(value)
			}
		} else {
			panic("parse error")
		}
	}

	/* Use unsecure password authorization */
	if streamCommandId == stream.M_ADR {
		if mailer.respAuthorization != "" {
			return NewMailerStateSecureAuthRemoteAction()
		} else {
			return NewMailerStateAuthRemote()
		}
	}

	return self
}

func (self *AdditionalStepAction) processFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {
	if nextFrame.Command {
		return self.processCommandFrame(mailer, nextFrame)
	} else {
		return NewMailerStateEnd()
	}
}

func (self *AdditionalStepAction) Process(mailer *Mailer) IMailerState {
	select {
	case nextFrame := <-mailer.stream.InDataFrames:
		return self.processFrame(mailer, nextFrame)
	}
	return self
}
