package mailer

import (
	"bytes"
	"fmt"
	cmn "github.com/vit1251/golden/pkg/common"
	"log"
	"os"
	"path"
	"time"
)

type MailerStateReceive struct {

}

func NewMailerStateReceive() *MailerStateReceive {
	msr := new(MailerStateReceive)
	return msr
}

func (self *MailerStateReceive) String() string {
	return "MailerStateReceive"
}

func (self *MailerStateReceive) Process(mailer *Mailer) IMailerState {

	select {
	case nextFrame := <- mailer.inDataFrames:
		if nextFrame.Command {
			if nextFrame.CommandFrame.CommandID == M_EOB {

				return NewMailerStateTransfer()

			} else if nextFrame.CommandFrame.CommandID == M_FILE {
			//
			if mailer.outStream != nil {
				mailer.outStream.Close()
				mailer.InFileCount += 1
			}
			//
			log.Printf("Start receive: row = %s", nextFrame.CommandFrame.Body)
			// p0018ea8.WE0 39678 1579714843 0
			parts := bytes.SplitN(nextFrame.CommandFrame.Body, []byte(" "), 4)
			//
			filename := string(parts[0])
			size, _ := cmn.ParseSize(parts[1])
			unixtime, _ := cmn.ParseSize(parts[2])
			offset, _ := cmn.ParseSize(parts[3])
			//
			if offset != 0 {
				panic("Wrong offset is not 0")
			}
			//
			mailer.recvName = filename
			mailer.size = size
			mailer.recvUnix = unixtime
			//
			newPath := path.Join(mailer.workInbound, filename)
			log.Printf("RX stream save in %s", newPath)
			mailer.workPath = newPath
			//
			if stream, err1 := os.Create(newPath); err1 != nil {
				panic(err1)
			} else {
				mailer.outStream = stream
			}
		}
	} else {
		log.Printf("Data frame: body = %d", len(nextFrame.DataFrame.Body))
		//
		mailer.writeSize += len(nextFrame.DataFrame.Body)
		//
		if mailer.outStream == nil {
			panic("Data section outside stream.")
		}
		//
		mailer.outStream.Write(nextFrame.DataFrame.Body)
		//
		if mailer.writeSize == mailer.size {
			log.Printf("!!! RX streame complete. Close write stream.")
			//
			mailer.outStream.Close()
			mailer.outStream = nil
			//
			rawComplete := fmt.Sprintf("%s %d %d", mailer.recvName, mailer.writeSize, mailer.recvUnix)
			mailer.writeCommandPacket(M_GOT, []byte(rawComplete))
			//
			newPath := path.Join(mailer.inboundDirectory, mailer.recvName)
			log.Printf("Rename %s -> %s", mailer.workPath, newPath)
			os.Rename(mailer.workPath, newPath)
			//
			mailer.writeSize = 0
		}
		//
	}

	case <-time.After(1 * time.Second):
		log.Printf("!!!Timeout!!!")
		return self
	}

	return self
}


