package mailer

import (
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
