package mailer

type RxState int

const (
	RxWaitF  RxState = 0x0101
	RxAccF   RxState = 0x0102
	RxRaceD  RxState = 0x0103
	RxWriteD RxState = 0x0104
	RxEOB    RxState = 0x0105
	RxDone   RxState = 0x0106
)

type TxState int

const (
	TxGNF   TxState = 0x0201
	TxTryR  TxState = 0x0202
	TxReadS TxState = 0x0203
	TxWLA   TxState = 0x0204
	TxDone  TxState = 0x0205
)

type FileTransferStageState int

const (
	FileTransferInitTransferState FileTransferStageState = 0
	FileTransferSwitchState       FileTransferStageState = 1
)
