package mailer

import (
	"bytes"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/mailer/cache"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
	"path"
)

type MailerStateRxWaitF struct {
	MailerState
}

func NewMailerStateRxWaitF() *MailerStateRxWaitF {
	return new(MailerStateRxWaitF)
}

func (self MailerStateRxWaitF) String() string {
	return "MailerStateRxWaitF"
}

func (self *MailerStateRxWaitF) Process(mailer *Mailer) IMailerState {

	/* Get a frame from Input Buffer */
	nextFrame := <-mailer.stream.InDataFrames
	mailer.stream.RemReadReady()

	/* Got Data frame */
	if nextFrame.IsDataFrame() {
		// ignore
		return NewMailerStateSwitch()
	}

	/* Got M_ERR */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_ERR {
			/* Report error */
			log.Printf("RxWaitF receive M_ERR")
			mailer.rxState = RxDone
			return NewMailerStateSwitch()
		}
	}

	/* Got M_GET / M_GOT / M_SKIP */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_GET || nextFrame.CommandID == stream.M_GOT || nextFrame.CommandID == stream.M_SKIP {
			/* Add frame to The Queue */
			// TODO --
			return NewMailerStateSwitch()
		}
	}

	/* Got M_EOB */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_EOB {
			mailer.rxState = RxEOB
			return NewMailerStateSwitch()
		}
	}

	/* Got M_FILE */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_FILE {
			self.processFile(mailer, nextFrame)
			mailer.rxState = RxAccF
			return NewMailerStateSwitch()
		}
	}

	/* Got other known frame */
	// TODO -

	/* Got unknown frame */
	// TODO -

	return NewMailerStateSwitch()

}

func (self MailerStateRxWaitF) processFile(mailer *Mailer, nextFrame stream.Frame) {

	packet := nextFrame.CommandFrame.Body

	log.Printf("Receive: row = %s", packet)

	/* Parsae incoming packet */
	// p0018ea8.WE0 39678 1579714843 0

	parts := bytes.SplitN(packet, []byte(" "), 4)

	recvName := string(parts[0])
	recvSize, _ := cmn.ParseSize(parts[1])
	recvUnixtime, _ := cmn.ParseSize(parts[2])
	//recvOffset, _ := cmn.ParseSize(parts[3])

	mailer.recvName = cache.NewFileEntry()
	mailer.recvName.AbsolutePath = path.Join(mailer.workInbound, recvName) // TODO - unescape file name ...
	mailer.recvName.Name = recvName

	mailer.readSize = int64(recvSize)
	mailer.recvUnix = recvUnixtime
	//mailer.recvOffset = ...

}
