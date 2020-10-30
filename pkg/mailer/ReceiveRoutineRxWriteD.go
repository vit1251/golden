package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"io"
	"log"
	"os"
	"path"
)


func makeFileGotPacket(mailer *Mailer, recvOffset int64) {

	recvName := mailer.recvName.Name
	recvUnix := mailer.recvUnix

	rawComplete := fmt.Sprintf("%s %d %d", recvName, recvOffset, recvUnix)
	mailer.stream.WriteCommandPacket(stream.M_GOT, []byte(rawComplete))

}

func ReceiveRoutineRxWriteD(mailer *Mailer) ReceiveRoutineResult {

	/* Write data to file */
	var err error = nil

	/* Write Failed */
	if err != nil {
		/* Report error */
		log.Printf("Report error")
		mailer.rxState = RxDone
		return RxFailure
	}

	offset, err1 := mailer.recvStream.Seek(0, io.SeekCurrent)
	if err1 != nil {
		log.Printf("Error: err %+v", err1)
		mailer.rxState = RxDone
		return RxFailure
	}

	log.Printf("RxWriteD: offset = %d expected = %d", offset, mailer.readSize)

	/* File Pos > Reported */
	if offset > mailer.readSize {
		/* Report write beyond EOF */
		log.Printf("Report write beyond EOF")
		mailer.rxState = RxDone
		return RxFailure
	}

	/* File Pos = Reported */
	if offset == mailer.readSize {

		/* Close File */
		mailer.recvStream.Close()
		mailer.recvStream = nil

		/* Send M_GOT */
		makeFileGotPacket(mailer, offset)

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
		return RxOk
	}

	if offset < mailer.readSize {
		mailer.rxState = RxRaceD
		return RxOk
	}

	panic("unknown case or memory corruption")

}
