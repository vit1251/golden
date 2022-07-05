package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"github.com/vit1251/golden/pkg/queue"
	"log"
	"os"
)

func popNextFileEntry(mailer *Mailer) *queue.FileEntry {

	var result *queue.FileEntry = nil

	queueSize := len(mailer.outboundQueue)
	if queueSize > 0 {
		result = &mailer.outboundQueue[0]
		mailer.outboundQueue = mailer.outboundQueue[1:]
	}

	return result

}

func makeFilePakcet(mailer *Mailer, stm *os.File) error {

	/* File summary */
	streamInfo, err1 := stm.Stat()
	if err1 != nil {
		return err1
	}

	/* Prepare M_FILE packet */
	// p0018ea8.WE0 39678 1579714843 0

	sendSize := streamInfo.Size()
	sendTime := streamInfo.ModTime().Unix()
	sendName := mailer.sendName.Name

	packet := fmt.Sprintf("%s %d %d %d", sendName, sendSize, sendTime, 0)

	mailer.stream.WriteHeader(packet)

	return nil

}

func TransmitRoutineTxGNF(mailer *Mailer) TransmitRoutineResult {

	/* Open next file from outgoing queue */
	mailer.sendName = popNextFileEntry(mailer)

	/* File opened OK */
	if mailer.sendName != nil {

		log.Printf("TX name: nextFile = %+v", mailer.sendName)

		stm, err1 := os.Open(mailer.sendName.AbsolutePath)

		if err1 == nil {

			mailer.sendStream = stm

			/* Send M_FILE */
			makeFilePakcet(mailer, stm)

			/* Report sending file */
			log.Printf("Start sending file - %+v", mailer.sendName)

			/* Next state */
			mailer.txState = TxTryR
			return TxContinue
		}

		/* Failed to open file */
		if err1 != nil {
			/* Report failure */
			log.Printf("Fail to open file")
			/* New state */
			mailer.txState = TxDone
			return TxFailure
		}

	}

	/* No more files */
	if mailer.sendName == nil {

		/* Send M_EOB */
		mailer.stream.WriteCommandPacket(stream.M_EOB, []byte("Complete!"))

		/* Report end of batch */
		log.Printf("Transmite Routine: End of Batch.")

		mailer.txState = TxWLA
		return TxContinue
	}

	panic("unknown case or memory corruption")

}
