package mailer

import (
	"io"
	"log"
)

func TransmitRoutineTxReadS(mailer *Mailer) TransmitRoutineResult {

	/* Read data block from file */
	chunkSize := 4096 // TODO - cmn.Min(1024, 4096)
	chunk := make([]byte, chunkSize)

	sendReady, err3 := mailer.sendStream.Read(chunk)
	if sendReady > 0 {
		chunk = chunk[:sendReady]
	} else {
		chunk = nil
	}

	/* Read failed */
	if err3 != nil {
		if err3 != io.EOF {
			/* Report error */
			log.Printf("Error reading TX stream: err = %+v", err3)
			mailer.txState = TxDone
			return TxFailure
		}
	}

	/* Read OK, Reaced EOF */
	if err3 == io.EOF {

		/* Send data block frame */
		if chunk != nil {
			mailer.stream.WriteData(chunk)
		}

		/* Close current file */
		mailer.sendStream.Close()
		mailer.sendStream = nil

		/* Add current file to Pending Files */
		if mailer.sendName != nil {
			mailer.pendingFiles.Push(*mailer.sendName)
		}

		/* Next state */
		mailer.txState = TxGNF
		return TxOk
	}

	/* Read OK, not reached EOF */
	if err3 == nil {

		/* Send data block frame */
		mailer.stream.WriteData(chunk)

		/* Next state */
		mailer.txState = TxTryR
		return TxOk
	}

	panic("unknown case or memory corruption")

}
