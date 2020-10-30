package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"io"
	"log"
	"os"
	"path"
)

type MailerStateRxWriteD struct {
	MailerState
}

func NewMailerStateRxWriteD() *MailerStateRxWriteD {
	return new(MailerStateRxWriteD)
}

func (self MailerStateRxWriteD) String() string {
	return "MailerStateRxWriteD"
}

func (self *MailerStateRxWriteD) Process(mailer *Mailer) IMailerState {


	offset, err := mailer.recvStream.Seek(0, io.SeekCurrent)
	if err != nil {
		log.Printf("Error: err %+v", err)
	}

	/* File Pos > Reported */
	if offset > mailer.readSize {
		// TODO - Report write beyond EOF
	}

	/* File Pos = Reported */
	if offset == mailer.readSize {

		/* Close File */
		mailer.recvStream.Close()
		mailer.recvStream = nil

		/* Send M_GOT */
		self.makeFileGotPacket(mailer, offset)

		/* Report File Received */
		log.Printf("Recieved file - %s", mailer.recvName)
		mailer.InFileCount += 1

		/* Move in incoming directory */
		newPath := path.Join(mailer.inboundDirectory, mailer.recvName.Name)
		log.Printf("Rename %s -> %s", mailer.recvName.AbsolutePath, newPath)
		if err1 := os.Rename(mailer.recvName.AbsolutePath, newPath); err1 != nil {
			log.Printf("Rename error!")
		}

		mailer.recvName = nil

		mailer.rxState = RxWaitF
	}

	if offset < mailer.readSize {
		mailer.rxState = RxRaceD
	}

	return NewMailerStateSwitch()

}

func (self MailerStateRxWriteD) makeFileGotPacket(mailer *Mailer, recvOffset int64) {

	recvName := mailer.recvName.Name
	recvUnix := mailer.recvUnix

	rawComplete := fmt.Sprintf("%s %d %d", recvName, recvOffset, recvUnix)
	mailer.stream.WriteCommandPacket(stream.M_GOT, []byte(rawComplete))

}
