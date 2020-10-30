package mailer

import (
	"io"
	"log"
	"os"
	"path"
)

type MailerStateTxReadS struct {
	MailerState
}

func NewMailerStateTxReadS() *MailerStateTxReadS {
	return new(MailerStateTxReadS)
}

func (self MailerStateTxReadS) String() string {
	return "MailerStateTxReadS"
}

func (self *MailerStateTxReadS) Process(mailer *Mailer) IMailerState {

	/* Read data block from file */
	chunkSize := 4096 // TODO - cmn.Min(1024, 4096)
	chunk := make([]byte, chunkSize)

	sendReady, err3 := mailer.sendStream.Read(chunk)

	/* Read failed */
	if err3 != nil {
		if err3 != io.EOF {
			/* Report error */
			log.Printf("Error reading TX stream: err = %+v", err3)
			mailer.txState = TxDone
		}
	}

	/* Read OK, Reaced EOF */
	if err3 == io.EOF {

		/* Send data block frame */
		if sendReady > 0 {
			mailer.stream.WriteData(chunk)
		}

		/* Close current file */
		mailer.sendStream.Close()
		mailer.sendStream = nil

		/* Complete routine */
		newName := path.Join(mailer.GetWorkOutbound(), mailer.sendName.Name)
		err4 := os.Rename(mailer.sendName.AbsolutePath, newName)
		if err4 != nil {
			log.Printf("Send file rename error: err = %+v")

		}

		/* Add current file to Panding Files */
		mailer.txState = TxGNF

	}

	/* Read OK, not reached EOF */
	if err3 == nil {

		/* Send data block frame */
		mailer.stream.WriteData(chunk)

		/* Next state */
		mailer.txState = TxTryR

	}

	return NewMailerStateSwitch()

}
