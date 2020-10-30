package stream

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	FrameCommandMask uint16 = 0x08000
	FrameSizeMask    uint16 = 0x07FFF
)

func (self *MailerStream) parseFrameHeader(frameHeader uint16) (bool, uint16) {
	var frameCommand bool = frameHeader &FrameCommandMask == FrameCommandMask
	var frameSize uint16 = frameHeader & FrameSizeMask
	return frameCommand, frameSize
}

func (self *MailerStream) pushPacket(nextFrame Frame) {
	for pending := true; pending; {
		select {

			case self.InFrameReady <- nil:
				log.Printf("MailerStream: RX stream: marker read READY push")

			case self.InFrame <- nextFrame:
				log.Printf("MailerStream: RX stream: ask packet with content")
				pending = false

		}
	}
}

func (self *MailerStream) processRX() {

	log.Printf("MailerStream: RX stream: start")
	for {

		log.Printf("vvv RX stream vvv")

		/* Receive frame header */
		var frameHeader uint16
		err1 := binary.Read(self.reader, binary.BigEndian, &frameHeader)
		if err1 != nil {
			log.Printf("MailerStream: Read: err = %+v", err1)
			break
		}
		log.Printf("RX frame: header %04X", frameHeader)

		/* Parse frame header */
		frameCommandFlag, frameSize := self.parseFrameHeader(frameHeader)
		log.Printf("RX frame: command %v size = %d", frameCommandFlag, frameSize)
		frameBody := make([]byte, frameSize)
		_, err2 := io.ReadFull(self.reader, frameBody)
		if err2 != nil {
			log.Printf("Frame RX body error.")
			break
		}

		/* Push frame in queue */
		if frameCommandFlag {
			var commandID CommandID = CommandID(frameBody[0])
			nextFrame := Frame{
				Command: true,
				CommandFrame: CommandFrame{
					CommandID: commandID,
					Body: frameBody[1:],
				},
			}

			log.Printf("MailerStream: RX frame: commandID = %q body = %s", nextFrame.CommandFrame.CommandID, nextFrame.CommandFrame.Body)
			self.pushPacket(nextFrame)

		} else {
			/* Store data frame in queue */
			nextFrame := Frame{
				Command: false,
				DataFrame: DataFrame{
					Body: frameBody,
				},
			}
			self.pushPacket(nextFrame)
		}
		log.Printf("^^^ RX stream ^^^")
	}

	log.Printf("MailerStream: RX stream: stop")

	/* Release resources */
	log.Printf("MailerStream: processRX: stream EOF")
	close(self.InFrameReady)
	close(self.InFrame)

	/* Done reader */
	self.wait.Done()

}
