package mailer

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type MailerStateTransfer struct {
	MailerState
}

func NewMailerStateTransfer() *MailerStateTransfer {
	mst := new(MailerStateTransfer)
	return mst
}

func (self *MailerStateTransfer) String() string {
	return "MailerStateTransfer"
}

func (self *MailerStateTransfer) Transmit(mailer *Mailer, i Item) error {

	/* Open stream */
	stream, err1 := os.Open(i.AbsolutePath)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

	/* Some status */
	streamInfo, err2 := stream.Stat()
	if err2 != nil {
		return err2
	}

	/* Transmit header */
	// p0018ea8.WE0 39678 1579714843 0
	streamSize := streamInfo.Size()
	streamTime := streamInfo.ModTime().Unix()
	fileStat := fmt.Sprintf("%s %d %d %d", i.Name, streamSize, streamTime, 0)
	log.Printf("TX %s", fileStat)
	mailer.writeHeader(fileStat)

	/* Transmit chunk */
	var outSize int = int(streamSize)
	for {

		/* Calculate transmit chunk */
		chunkSize := Min(outSize, 4096)
		chunk := make([]byte, chunkSize)

		/* Read */
		_, err3 := io.ReadFull(stream, chunk)
		if err3 != nil {
			return err3
		}

		/* Transmit chunk */
		mailer.writeData(chunk)

		/* Update TX size */
		outSize -= chunkSize
		if outSize == 0 {
			log.Printf("Transmit complete!")
			mailer.OutFileCount += 1
			break
		}
	}

	/* Check error */

	return nil
}

func (self *MailerStateTransfer) Process(mailer *Mailer) IMailerState {

	mo := NewMailerOutbound(mailer.SetupManager)
	items, err2 := mo.GetItems()
	if err2 != nil {
		panic(err2)
	}

	for _, item := range items {

		/* Transmit packet */
		err1 := self.Transmit(mailer, *item)
		if err1 != nil {
			log.Printf("Unable transmit %s file: err = %+v", item.Name, err1)
			break
		}

		/* Complete routine */
		newName := path.Join(mailer.TempOutbound, item.Name)
		os.Rename(item.AbsolutePath, newName)

	}

	log.Printf("Sent complete!")

	return NewMailerStateCloseConnection()
}
