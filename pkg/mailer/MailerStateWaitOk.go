package mailer

import (
	"fmt"
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

func (self *MailerStateWaitOk) processCommandFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {
	command := nextFrame.CommandFrame.CommandID

	if command == stream.M_NUL {
		log.Printf("Warning: ... NUL packet during authorization state ...")
		return self
	} else if command == stream.M_OK {
		log.Printf("Auth - OK")
		return NewMailerStateOpts()
	} else if command == stream.M_ERR {
		log.Printf("AUTH - ERROR: err = %+v", nextFrame.CommandFrame.Body)
		mailer.report.SetStatus(fmt.Sprintf("Authorization error: reason = %+v", nextFrame.CommandFrame.Body))
	}

	return nil
}

func (self *MailerStateWaitOk) processFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {

	if nextFrame.IsCommandFrame() {
		return self.processCommandFrame(mailer, nextFrame)
	} else {
		log.Printf("Unexpected frame: frame = %+v", nextFrame)
	}

	return self
}

func (self *MailerStateWaitOk) Process(mailer *Mailer) IMailerState {

	select {

	case nextFrame := <-mailer.stream.InFrame:
		return self.processFrame(mailer, nextFrame)

		//	case <-mailer.WaitAddrTimeout:
		//		log.Printf("Timeout!")

	}

	return self

}
