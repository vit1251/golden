package mailer

import "log"

type MailerStateSecure struct {

}

func NewMailerStateSecure() *MailerStateSecure {
	mss := new(MailerStateSecure)
	return mss
}

func (self *MailerStateSecure) String() string {
	return "MailerStateSecure"
}

func (self *MailerStateSecure) Process(mailer *Mailer) IMailerState {

	nextFrame := <-mailer.inDataFrames

	log.Printf("Auth complete: frame = %+v", nextFrame)

	if nextFrame.Command {
		if nextFrame.CommandFrame.CommandID == M_OK {
			return NewMailerStateReceive()
		}
	} else {
		log.Panicf("Unexpected frame")
	}

	return self
}