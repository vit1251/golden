package stream

import (
	"encoding/binary"
	"fmt"
	"log"
)

func (self *MailerStream) writeCommentPacket(msg []byte) (error) {
	return self.WriteCommandPacket(M_NUL, msg)
}

func (self *MailerStream) WriteInfo(name string, value string) (error) {
	var row string = fmt.Sprintf("%s %s", name, value)
	var commentPacket1 []byte = []byte(row)
	return self.writeCommentPacket(commentPacket1)
}

func (self *MailerStream) WriteComment(comment string) (error) {
    var commentPacket1 []byte = []byte(comment)
    return self.writeCommentPacket(commentPacket1)
}

func (self *MailerStream) WriteAddress(addr string) (error) {
	var addrPacket []byte = []byte(addr)
	return self.WriteCommandPacket(M_ADR, addrPacket)
}

func (self *MailerStream) WritePassword(password string) (error) {
	var passwordPacket []byte = []byte(password)
	return self.WriteCommandPacket(M_PWD, passwordPacket)
}

func (self *MailerStream) processTXpacket(nextFrame Frame) {
	log.Printf("MailerStream: TX stream process frame")
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
		log.Printf("MailerStream: TX frame: commandID = %q", commandID)
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

func (self *MailerStream) processTX() {

	log.Printf("MailerStream: TX stream: start")
	for nextFrame := range self.OutDataFrames {
		log.Printf("TX packet processing")
		self.processTXpacket(nextFrame)
	}
	log.Printf("MailerStream: TX stream: stop")

	/* Done writer */
	self.wait.Done()

}
