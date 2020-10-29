package mailer

import (
	"bytes"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type MailerStateWaitAddrAction struct {
	IMailerState
}

func NewMailerStateWaitAddrAction() *MailerStateWaitAddrAction {
	msrh := new(MailerStateWaitAddrAction)
	return msrh
}

func (self *MailerStateWaitAddrAction) String() string {
	return "MailerStateWaitAddrAction"
}

func (self *MailerStateWaitAddrAction) processCommandFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {

	var streamCommandId = nextFrame.CommandFrame.CommandID

	if streamCommandId == stream.M_NUL {
		if key, value, err1 := mailer.parseInfoFrame(nextFrame.CommandFrame.Body); err1 == nil {
			log.Printf("Remote side option: name = %s value = %s", key, value)
			if bytes.Equal(key, []byte("OPT")) {
				mailer.parseInfoOptFrame(value)
			}
			return self
		} else {
			panic("parse error")
		}
	} else if streamCommandId == stream.M_BSY {
		panic("busy remote system")
	} else {
		log.Printf("Unexpected Frame in state: %v", mailer.sessionSetupState)
	}

	return nil
}

func (self *MailerStateWaitAddrAction) processFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {
	if nextFrame.Command {
		return self.processCommandFrame(mailer, nextFrame)
	} else {
		return NewMailerStateCloseConnection()
	}
}

func (self *MailerStateWaitAddrAction) Process(mailer *Mailer) IMailerState {

	select {

	case nextFrame := <-mailer.stream.InDataFrames:
		return self.processFrame(mailer, nextFrame)

//	case <-mailer.WaitAddrTimeout:
//		log.Printf("Timeout!")

	}

	return self

}
