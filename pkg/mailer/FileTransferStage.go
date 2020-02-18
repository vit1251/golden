package mailer

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

type RxState int

const (
	RxWaitF RxState = 0x0101
	RxAccF RxState = 0x0102
	RxRaceD RxState = 0x0103
	RxWriteD RxState = 0x0104
	RxEOB RxState = 0x0105
	RxDone RxState = 0x0106
)

type TxState int

const (
	TxGNF TxState = 0x0201
	TxTryR TxState = 0x0202
	TxReadS TxState = 0x0203
	TxWLA TxState = 0x0204
	TxDone TxState = 0x0205
)

type FileTransferStageState int

const (
	FileTransferInitTransferState FileTransferStageState = 0
	FileTransferSwitchState FileTransferStageState = 1
)

func (self *Mailer) action_T0_InitTransfer() {

	/* Set Timer */

	/* SetRxState to RxWaitF */

	/* Set TxState to TxGNF */

}

func (self *Mailer) processReceiveRoutine() {
}

func (self *Mailer) processTransmitRoutine() {
}

func (self *Mailer) action_T1_Switch() {

	/* Predicate 1: RxStat == RxDone and TxState is TxDone */
	if self.rxState == RxDone && self.txState == TxDone {
		//self.SetFileTransferState(nil)
	} else

	/* Predicate 2: Data available in Input Buffer */
	if len(self.inDataFrames) > 0 {
		self.processReceiveRoutine()
	} else

	/* Predicate 3: Free space exists in iutput buffer */
	if len(self.outDataFrames) < cap(self.outDataFrames) {
		self.processTransmitRoutine()
	}

	/* Predicate 4: Nothing happend */

	/* Predicate 5. Timer Expired */

}

func (self *Mailer) parseSize(size []byte) (int) {
	var num string = string(size)
	newSize, err1 := strconv.ParseInt(num, 10, 32)
	if err1 != nil {
		panic(err1)
	}
	return int(newSize)
}

func (self *Mailer) processReceiveFileTransferState() error {
	nextFrame := <- self.inDataFrames
	if nextFrame.Command {
		if nextFrame.CommandFrame.CommandID == M_FILE {
			//
			if self.outStream != nil {
				self.outStream.Close()
			}
			//
			log.Printf("Start receive: row = %s", nextFrame.CommandFrame.Body)
			// p0018ea8.WE0 39678 1579714843 0
			parts := bytes.SplitN(nextFrame.CommandFrame.Body, []byte(" "), 4)
			//
			filename := string(parts[0])
			size := self.parseSize(parts[1])
			unixtime := self.parseSize(parts[2])
			offset := self.parseSize(parts[3])
			//
			if offset != 0 {
				panic("Wrong offset is not 0")
			}
			//
			self.recvName = filename
			self.size = size
			self.recvUnix = unixtime
			//
			newPath := path.Join(self.inboundDirectory, filename)
			log.Printf("RX stream save in %s", newPath)
			//
			if stream, err1 := os.Create(newPath); err1 != nil {
				panic(err1)
			} else {
				self.outStream = stream
			}
		}
	} else {
		log.Printf("Data frame: body = %d", len(nextFrame.DataFrame.Body))
		//
		self.writeSize += len(nextFrame.DataFrame.Body)
		//
		if self.outStream == nil {
			panic("Data section outside stream.")
		}
		//
		self.outStream.Write(nextFrame.DataFrame.Body)
		//
		if self.writeSize == self.size {
			log.Printf("!!! RX streame complete. Close write stream.")
			//
			self.outStream.Close()
			self.outStream = nil
			//
			rawComplete := fmt.Sprintf("%s %d %d", self.recvName, self.writeSize, self.recvUnix)
			self.writeCommandPacket(M_GOT, []byte(rawComplete))
			//
			self.writeSize = 0
		}
		//
	}
	//
	log.Printf("Receive complete!")
	//self.Tosser
	//
	return nil
}

func (self *Mailer) processTransmitFileTransferState() error {

	mo := NewMailerOutbound(self.SetupManager)
	items, err2 := mo.GetItems()
	if err2 != nil {
		panic(err2)
	}

	for _, item := range items {

		/* Transmit packet */
		err1 := self.Transmit(*item)
		if err1 != nil {
			log.Printf("Unable transmit %s file: err = %+v", item.Name, err1)
			break
		}

		/* Complete routine */
		newName := path.Join(self.TempOutbound, item.Name)
		os.Rename(item.AbsolutePath, newName)

	}

	log.Printf("Sent complete!")

	return nil
}

func (self *Mailer) processFileTransferState() error {
	if self.transferState == FileTransferInitTransferState {
		self.SetFileTransferState(FileTransferSwitchState)
	} else if self.transferState == FileTransferSwitchState {
		self.processReceiveFileTransferState()
		self.processTransmitFileTransferState()
		//
		// TODO - on receive and sent complete close stream and start toss ...
	} else {
		// TODO - error here ...
	}
	return nil
}
