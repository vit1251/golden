package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type MailerStateWaitOk struct {
	MailerState
}

func NewMailerStateWaitOk() *MailerStateWaitOk {
	return new(MailerStateWaitOk)
}

func (self *MailerStateWaitOk) String() string {
	return "MailerStateWaitOk"
}

func (self *MailerStateWaitOk) processCommandFrame(nextFrame stream.Frame) IMailerState {
	command := nextFrame.CommandFrame.CommandID

	if command == stream.M_NUL {
		log.Printf("...")
		return self
	} else if command == stream.M_OK {
		log.Printf("Auth - OK")
		return NewMailerStateOpts()
	} else if command == stream.M_ERR {
		log.Printf("AUTH - ERROR: err = %+v", nextFrame.CommandFrame.Body)
	}

	return nil
}

func (self *MailerStateWaitOk) processFrame(nextFrame stream.Frame) IMailerState {

	if nextFrame.IsCommandFrame() {
		return self.processCommandFrame(nextFrame)
	} else {
		log.Printf("Unexpected frame: frame = %+v", nextFrame)
	}

	return self
}

func (self *MailerStateWaitOk) Process(mailer *Mailer) IMailerState {

	select {

	case nextFrame := <-mailer.stream.InDataFrames:
		return self.processFrame(nextFrame)

		//	case <-mailer.WaitAddrTimeout:
		//		log.Printf("Timeout!")

	}

	return self

}
