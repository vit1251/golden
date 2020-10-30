package mailer

import (
	"bytes"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
	"os"
	"path"
)

func processGotPacket(mailer *Mailer, nextFrame stream.Frame) {

	log.Printf("Process M_GOT incoming packet")

	packet := nextFrame.CommandFrame.Body

	// abc.txt 000 123
	values := bytes.SplitN(packet, []byte(" "), 2)
	if len(values) > 1 {
		name := values[0]
		newName := string(name) // TODO - unescape name here ...

		log.Printf("ACK receive packet: name = %+v", name)
		mailer.pendingFiles.RemoveByName(newName)

		pendingPath := path.Join(mailer.outboundDirectory, newName)
		newPath := path.Join(mailer.workOutbound, newName)

		log.Printf("SEND COMPLETE: MOVE: %s -> %s", pendingPath, newPath)

		err4 := os.Rename(pendingPath, newPath)
		if err4 != nil {
			log.Printf("Send file rename error: err = %+v", err4)
		}

	} else {
		log.Printf("M_GOT parse error!")
	}

}

func ProcessTheQueue(mailer *Mailer) {

	log.Printf("ProcessTheQueue!!!")

	nextFrame := mailer.queue.Pop()

	/* M_GET received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_GET {
			// TODO - ...
		}
	}

	/* M_GET received for a known file */
	// TODO - ...

	/* M_GOT file that is currently transmitting */
	// TODO - ...

	/* M_GOT file that is not currently transmitting */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_GOT {
			processGotPacket(mailer, *nextFrame)
		}
	}


}
