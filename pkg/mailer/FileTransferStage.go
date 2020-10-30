package mailer

type RxState string

const (
	RxWaitF  RxState = "RxWaitF"
	RxAccF   RxState = "RxAccF"
	RxRaceD  RxState = "RxRaceD"
	RxWriteD RxState = "RxWriteD"
	RxEOB    RxState = "RxEOB"
	RxDone   RxState = "RxDone"
)

type TxState string

const (
	TxGNF   TxState = "TxGNF"
	TxTryR  TxState = "TxTryR"
	TxReadS TxState = "TxReadS"
	TxWLA   TxState = "TxWLA"
	TxDone  TxState = "TxDone"
)
