package mailer

import (
	"encoding/binary"
	"fmt"
	"log"
)

func (self *Mailer) writePacket(msg []byte) {

}

func (self *Mailer) writeCommandPacket(commandID CommandID, msgBody []byte) error {
	log.Printf("writeCommandPacket: commandID = %d body = %s", commandID, msgBody)
	newFrame := Frame{
		Command: true,
		CommandFrame: CommandFrame{
			CommandID: commandID,
			Body: msgBody,
		},
	}
	self.outDataFrames <- newFrame
	return nil
}

func (self *Mailer) writeCommentPacket(msg []byte) (error) {
	return self.writeCommandPacket(M_NUL, msg)
}

func (self *Mailer) WriteInfo(name string, value string) (error) {
	var row string = fmt.Sprintf("%s %s", name, value)
	var commentPacket1 []byte = []byte(row)
	return self.writeCommentPacket(commentPacket1)
}

func (self *Mailer) WriteComment(comment string) (error) {
    var commentPacket1 []byte = []byte(comment)
    return self.writeCommentPacket(commentPacket1)
}

func (self *Mailer) WriteAddress(addr string) (error) {
	var addrPacket []byte = []byte(addr)
	return self.writeCommandPacket(M_ADR, addrPacket)
}

func (self *Mailer) WritePassword(password string) (error) {
	var passwordPacket []byte = []byte(password)
	return self.writeCommandPacket(M_PWD, passwordPacket)
}

func (self *Mailer) processTXpacket(nextFrame Frame) {
	log.Printf("TX stream process frame")
	if nextFrame.Command {
		var frameHeader uint16
		var frameSize int = len(nextFrame.CommandFrame.Body)
		log.Printf("TX frame: type = command (%v): size = %d", nextFrame.Command, frameSize)
		if frameSize >= int(FrameSizeMask) {
			panic("Frame size is overflow.")
		}
		frameHeader = FrameCommandMask + uint16(frameSize) + 1
		log.Printf("TX frame: header %04X", frameHeader)
		binary.Write(self.writer, binary.BigEndian, frameHeader)
		var commandID uint8 = uint8(nextFrame.CommandFrame.CommandID)
		log.Printf("TX frame: commandID = %d", commandID)
		binary.Write(self.writer, binary.BigEndian, commandID)
		binary.Write(self.writer, binary.BigEndian, nextFrame.CommandFrame.Body)
		self.writer.Flush()
	} else {
		var frameHeader uint16
		var frameSize int = len(nextFrame.DataFrame.Body)
		log.Printf("TX frame: type = data: size = %d", frameSize)
		if frameSize >= int(FrameSizeMask) {
			panic("Frame size is overflow.")
		}
		frameHeader = uint16(frameSize)
		binary.Write(self.writer, binary.BigEndian, frameHeader)
		binary.Write(self.writer, binary.BigEndian, nextFrame.DataFrame.Body)
		self.writer.Flush()
	}
}

func (self *Mailer) processTX() {

	log.Printf("TX stream start")
	for alive := true; alive; {
		select {

		case nextFrame := <-self.outDataFrames:
			self.processTXpacket(nextFrame)
			break

		case <-self.outStop:
			alive = false
			break

		}
	}
	log.Printf("TX stream stop")

	/* Done waiter */
	self.wait.Done()

}