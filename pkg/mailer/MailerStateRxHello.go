package mailer

import (
	"bytes"
	"log"
)

type MailerStateRxHello struct {
	IMailerState
}

func NewMailerStateRxHello() *MailerStateRxHello {
	msrh := new(MailerStateRxHello)
	return msrh
}

func (self *MailerStateRxHello) String() string {
	return "MailerStateRxHello"
}

func (self *MailerStateRxHello) Process(mailer *Mailer) IMailerState {

	nextFrame := <-mailer.inDataFrames
	if nextFrame.Command {
		if nextFrame.CommandFrame.CommandID == M_NUL {
			key, value, err1 := mailer.parseInfoFrame(nextFrame.CommandFrame.Body)
			if err1 != nil {
				panic(err1)
			}
			log.Printf("Remote side option: name = %s value = %s", key, value)
			if bytes.Equal(key, []byte("OPT")) {
				mailer.parseInfoOptFrame(value)
			}
		} else if nextFrame.CommandFrame.CommandID == M_ADR {
			// TODO - save remote address in variable ...
			return NewMailerStateAuthorization()
		} else {
			log.Panicf("Unexpected Frame in state: %v", mailer.sessionSetupState)
		}
	}

	return self
}
