package mailer

import (
	"bytes"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type MailerStateWaitAddr struct {
	MailerState
}

func NewMailerStateWaitAddr() *MailerStateWaitAddr {
	msrh := new(MailerStateWaitAddr)
	return msrh
}

func (self *MailerStateWaitAddr) String() string {
	return "MailerStateWaitAddr"
}

func (self *MailerStateWaitAddr) processFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {

	/* M_ADR frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_ADR {
			return NewMailerStateAuthRemote()
		}
	}

	/* M_BSY frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_BSY {
			log.Printf("Remote system is BUSY")
			return NewMailerStateEnd()
		}
	}

	/* M_ERR frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_BSY {
			log.Printf("Remote system is ERROR")
			return NewMailerStateEnd()
		}
	}

	/* M_NUL frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_NUL {
			if key, value, err1 := mailer.parseInfoFrame(nextFrame.CommandFrame.Body); err1 == nil {
				log.Printf("Remote side option: name = %s value = %s", key, value)
				if bytes.Equal(key, []byte("OPT")) {
					mailer.parseInfoOptFrame(value)
				}
				return NewMailerStateWaitAddr()
			} else {
				log.Printf("M_NUL in WaitAddr parse error: err = %+v", err1)
			}
		}
	}

	return NewMailerStateWaitAddr()
}

func (self *MailerStateWaitAddr) Process(mailer *Mailer) IMailerState {

	nextFrame := mailer.stream.GetFrame()
	return self.processFrame(mailer, nextFrame)

}
