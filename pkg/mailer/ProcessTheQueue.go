package mailer

import (
	"bytes"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
	"os"
	"path"
)

type GotStat struct {
	Name     string
	Size     int64
	UnixTime int64
}

func parseGot(packet []byte) *GotStat {
	values := bytes.SplitN(packet, []byte(" "), 2)
	if len(values) > 1 {
		name := values[0]
		// TODO - Size ..
		// TODO - UnixTime ..
		result := GotStat{
			Name: string(name),      // TODO - unescape name here ...
		}
		return &result
	} else {
		return nil
	}
}


func processGotPacket(mailer *Mailer, nextFrame stream.Frame) {

	log.Printf("Process M_GOT incoming packet")
	packet := nextFrame.CommandFrame.Body
	gs := parseGot(packet)
	if gs == nil {
		log.Printf("M_GOT parse error!")
		return
	}

	/* M_GOT file that is currently transmitting */
	// TODO - ...

	/* M_GOT file that is not currently transmitting */
	if (mailer.sendName == nil) || (mailer.sendName != nil && mailer.sendName.Name != gs.Name) {

		/* Files is in PendingFiles list */
		if mailer.pendingFiles.Contains(gs.Name) {

			pendingPath := path.Join(mailer.outboundDirectory, gs.Name)
			newPath := path.Join(mailer.workOutbound, gs.Name)

			log.Printf("SEND COMPLETE: MOVE: %s -> %s", pendingPath, newPath)

			err4 := os.Rename(pendingPath, newPath)
			if err4 != nil {
				log.Printf("Send file rename error: err = %+v", err4)
			}

			/* Remove file from the PendingFiles list */
			log.Printf("ACK receive packet: name = %+v", gs.Name)
			mailer.pendingFiles.RemoveByName(gs.Name)

			/* Set TxState to TxGNF */
			mailer.txState = TxGNF

		} else {

			/* Ignore frame */

		}

	}

}

func ProcessTheQueue(mailer *Mailer) {

	log.Printf("ProcessTheQueue!!!")

	nextFrame := mailer.queue.Pop()

	/* M_GET received */
	/* M_GET received for a known file */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_GET {
			// TODO - ...
		}
	}

	/* M_GOT file that is currently transmitting */
	/* M_GOT file that is not currently transmitting */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_GOT {
			processGotPacket(mailer, *nextFrame)
		}
	}

	/* M_SKIP */
	// TODO -

	/* M_NUL */
	// TODO -

}
