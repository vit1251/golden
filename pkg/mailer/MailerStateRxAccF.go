package mailer

import (
	stream2 "github.com/vit1251/golden/pkg/mailer/stream"
	"log"
	"os"
)

type MailerStateRxAccF struct {
	MailerState
}

func NewMailerStateRxAccF() *MailerStateRxAccF {
	return new(MailerStateRxAccF)
}

func (self MailerStateRxAccF) String() string {
	return "MailerStateRxAccF"
}

func (self *MailerStateRxAccF) Process(mailer *Mailer) IMailerState {

	/* Accept from beginning */

	log.Printf("Open path: %+v", mailer.recvName.AbsolutePath)

	stream, err1 := os.Create(mailer.recvName.AbsolutePath)

	if err1 != nil {
		log.Printf("Fail to create file")
		// TODO - error report and packet ...
		mailer.stream.WriteCommandPacket(stream2.M_ERR, []byte("Unable to open file!"))
		mailer.rxState = RxDone
	}

	if err1 == nil {
		log.Printf("Start receivnig file")
		mailer.recvStream = stream
		mailer.rxState = RxRaceD
	}

	return NewMailerStateSwitch()

}
