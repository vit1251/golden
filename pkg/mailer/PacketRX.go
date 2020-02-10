package mailer

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	FrameCommandMask uint16 = 0x08000
	FrameSizeMask    uint16 = 0x07FFF
)

func (self *Mailer) parseFrameHeader(frameHeader uint16) (bool, uint16) {
	var frameCommand bool = frameHeader & FrameCommandMask == FrameCommandMask
	var frameSize uint16 = frameHeader & FrameSizeMask
	return frameCommand, frameSize
}

func (self *Mailer) processRX() {
	log.Printf("RX stream start")
	for {
		log.Printf("vvv RX stream vvv")

		/* Receive frame header */
		var frameHeader uint16
		err1 := binary.Read(self.reader, binary.BigEndian, &frameHeader)
		if err1 == io.EOF {
			log.Printf("Session close.")
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
			log.Printf("RX frame: commandID = %d", commandID)
			log.Printf("RX frame: body = %s", frameBody[1:])
			nextFrame := Frame{
				Command: true,
				CommandFrame: CommandFrame{
					CommandID: commandID,
					Body: frameBody[1:],
				},
			}
			self.inDataFrames <- nextFrame
		} else {
			/* Store data frame in queue */
			nextFrame := Frame{
				Command: false,
				DataFrame: DataFrame{
					Body: frameBody,
				},
			}
			self.inDataFrames <- nextFrame
		}
		log.Printf("^^^ RX stream ^^^")
	}
	log.Printf("RX stream stop")
}
