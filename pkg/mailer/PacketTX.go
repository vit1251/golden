package mailer

import (
	"log"
	"fmt"
	"errors"
	"encoding/binary"
)

func (self *Mailer) writePacket(msg []byte) {

}

func (self *Mailer) writeCommandPacket(messageType MessageType, msg []byte) (error) {

	var packetHeader uint16
	var packetSize int = len(msg)
	log.Printf("Prepare packet: payload = %d", packetSize)
	if packetSize >= int(FrameSizeMask)  {
		return errors.New("Too big packet.")
	}
	packetHeader = FrameCommandMask + uint16(packetSize) + 1
	binary.Write(self.writer, binary.BigEndian, packetHeader)
	var operation uint8 = uint8(messageType)
	binary.Write(self.writer, binary.BigEndian, operation)
	binary.Write(self.writer, binary.BigEndian, msg)
	self.writer.Flush()
	return nil
}

func (self *Mailer) writeCommentPacket(msg []byte) (error) {
	return self.writeCommandPacket(TextFrame, msg)
}

func (self *Mailer) writeInfo(name string, value string) (error) {
	var row string = fmt.Sprintf("%s %s", name, value)
	var commentPacket1 []byte = []byte(row)
	return self.writeCommentPacket(commentPacket1)
}

func (self *Mailer) writeComment(comment string) (error) {
    var commentPacket1 []byte = []byte(comment)
    return self.writeCommentPacket(commentPacket1)
}

func (self *Mailer) writeAddress(addr string) (error) {
	var addrPacket []byte = []byte(addr)
	return self.writeCommandPacket(AddrFrame, addrPacket)
}

func (self *Mailer) writePassword(password string) (error) {
	var passwordPacket []byte = []byte(password)
	return self.writeCommandPacket(PasswordFrame, passwordPacket)
}
