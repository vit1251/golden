package mailer

import (
	"bytes"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"github.com/vit1251/golden/pkg/queue"
	"log"
	"path"
)

func processFile(mailer *Mailer, nextFrame stream.Frame) {

	packet := nextFrame.CommandFrame.Body

	log.Printf("Receive: row = %s", packet)

	/* Parsae incoming packet */
	// p0018ea8.WE0 39678 1579714843 0

	parts := bytes.SplitN(packet, []byte(" "), 4)

	recvName := string(parts[0])
	recvSize, _ := cmn.ParseSize(parts[1])
	recvUnixtime, _ := cmn.ParseSize(parts[2])
	//recvOffset, _ := cmn.ParseSize(parts[3])

	mailer.recvName = queue.NewFileEntry()
	mailer.recvName.AbsolutePath = path.Join(mailer.workInbound, recvName) // TODO - unescape file name ...
	mailer.recvName.Name = recvName

	mailer.readSize = int64(recvSize)
	mailer.recvUnix = recvUnixtime
	//mailer.recvOffset = ...

}

func ReceiveRoutineRxWaitF(mailer *Mailer) ReceiveRoutineResult {

	/* Get a frame from Input Buffer */
	nextFrame := <-mailer.stream.InFrame

	/* Got Data frame */
	if nextFrame.IsDataFrame() {
		/* ignore */
		return RxOk
	}

	/* Got M_ERR */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_ERR {
			/* Report error */
			log.Printf("RxWaitF receive M_ERR")
			mailer.rxState = RxDone
			return RxFailure
		}
	}

	/* Got M_GET / M_GOT / M_SKIP */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_GET || nextFrame.CommandFrame.CommandID == stream.M_GOT || nextFrame.CommandFrame.CommandID == stream.M_SKIP {
			/* Add frame to The Queue */
			log.Printf("State: RxWaitF TheQueue push packet")
			mailer.queue.Push(nextFrame)
			mailer.queue.Dump()
			return RxOk
		}
	}

	/* Got M_NUL */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_EOB {
			return RxOk
		}
	}

	/* Got M_EOB */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_EOB {
			mailer.rxState = RxEOB
			return RxOk
		}
	}

	/* Got M_FILE */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandID == stream.M_FILE {
			processFile(mailer, nextFrame)
			mailer.rxState = RxAccF
			return RxContinue
		}
	}

	/* Got other known frame */
	if nextFrame.IsCommandFrame() { // TODO - Other known frame ...
		/* Report unexpected frame */
		log.Printf("Report unexpected frame")
		mailer.rxState = RxDone
		return RxFailure
	}

	/* Got unknown frame */
	return RxOk

}
