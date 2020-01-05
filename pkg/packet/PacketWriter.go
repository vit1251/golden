package packet

import (
	"bufio"
	"os"
)

type PacketWriter struct {
	stream                   *os.File      /* Native OS stream */
	streamWriter             *bufio.Writer /* Cache */
	binaryStreamWriter       *BinaryWriter /* ...              */
}

func NewPacketWriter(name string) (*PacketWriter, error) {

	/* Crete new packet writer */
	pw := new(PacketWriter)

	/* Create native OS stream */
	if stream, err := os.Create(name); err != nil {
		return nil, err
	} else {
		pw.stream = stream
	}

	/* Create cache stream */
	pw.streamWriter = bufio.NewWriter(pw.stream)

	/* Create binary stream reader */
	if binaryStreamWriter, err := NewBinaryWriter(pw.streamWriter); err != nil {
		pw.Close()
		return nil, err
	} else {
		pw.binaryStreamWriter = binaryStreamWriter
	}

	/* Done */
	return pw, nil
}

func (self *PacketWriter) writePacketHeaderDateTime(pktDateTime *PktDateTime) (error) {

	/* Writing */
	if err1 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Year); err1 != nil {
		return err1
	}
	if err2 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Mon); err2 != nil {
		return err2
	}
	if err3 := self.binaryStreamWriter.WriteUINT16(pktDateTime.MDay); err3 != nil {
		return err3
	}
	if err4 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Hour); err4 != nil {
		return err4
	}
	if err5 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Min); err5 != nil {
		return err5
	}
	if err6 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Sec); err6 != nil {
		return err6
	}

	return nil
}

const PRODUCT_VERSION_MAJOR = 1
const PRODUCT_VERSION_MINOR = 2

func (self *PacketWriter) writePacketHeaderProductVersion() (error) {

	if err1 := self.binaryStreamWriter.WriteUINT8(PRODUCT_VERSION_MAJOR); err1 != nil {
		return err1
	}

	if err2 := self.binaryStreamWriter.WriteUINT8(PRODUCT_VERSION_MINOR); err2 != nil {
		return err2
	}

	return nil
}

func (self *PacketWriter) writePacketHeaderCapatiblityBytes(capByte1 uint8, capByte2 uint8) (error) {

	if err1 := self.binaryStreamWriter.WriteUINT8(capByte1); err1 != nil {
		return err1
	}

	if err2 := self.binaryStreamWriter.WriteUINT8(capByte2); err2 != nil {
		return err2
	}

	return nil
}

const PKT_VERSION = 2

func (self *PacketWriter) WritePacketHeader(pktHeader *PacketHeader) (error) {

	/* Write orginator node address */
	if err1 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigAddr.Node); err1 != nil {
		return err1
	}
	/* Write destination node address */
	if err2 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestAddr.Node); err2 != nil {
		return err2
	}

	/* Write packet create (12 byte) */
	if err3 := self.writePacketHeaderDateTime(&pktHeader.pktCreated); err3 != nil {
		return err3
	}

	/* Write unused (2 byte) */
	if err4 := self.binaryStreamWriter.WriteUINT16(0); err4 != nil {
		return err4
	}

	/* Write packet version (2 byte) */
	if err5 := self.binaryStreamWriter.WriteUINT16(PKT_VERSION); err5 != nil {
		return err5
	}

	/* Write origination network (2 byte) */
	if err6 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigAddr.Net); err6 != nil {
		return err6
	}

	/* Write destination network (2 byte) */
	if err7 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestAddr.Net); err7 != nil {
		return err7
	}

	/* Write product version (2 byte)*/
	if err8 := self.writePacketHeaderProductVersion(); err8 != nil {
		return err8
	}

	/* Write packet password (8 byte) */
	if err10 := self.binaryStreamWriter.WriteBytes(pktHeader.pktPassword); err10 != nil {
		return err10
	}

	/* Write packet zone (2 byte) */
	if err11 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigAddr.Zone); err11 != nil {
		return err11
	}

	/* Write packet zone (2 byte) */
	if err12 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestAddr.Zone); err12 != nil {
		return err12
	}

	/* Write auxNet (2 byte) */
	if err13 := self.binaryStreamWriter.WriteUINT16(pktHeader.auxNet); err13 != nil {
		return err13
	}

	/* Write capatiblity bytes (2 byte) */
	if err14 := self.writePacketHeaderCapatiblityBytes(0, 1); err14 != nil {
		return err14
	}

	/* Write product version (2 byte) */
	if err15 := self.writePacketHeaderProductVersion(); err15 != nil {
		return err15
	}

	/* Write capability word (2 byte) */
	if err16 := self.binaryStreamWriter.WriteUINT16(1); err16 != nil {
		return err16
	}

	/* Write additional zone info (2 byte) */
	if err19 := self.binaryStreamWriter.WriteUINT16(0); err19 != nil {
		return err19
	}
	/* Write additional zone info (2 byte) */
	if err20 := self.binaryStreamWriter.WriteUINT16(0); err20 != nil {
		return err20
	}

	/* Write point (2 byte) */
	if err21 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigAddr.Point); err21 != nil {
		return err21
	}

	/* Write point (2 byte) */
	if err22 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestAddr.Point); err22 != nil {
		return err22
	}

	/* Write production data (2 byte) */
	if err23 := self.binaryStreamWriter.WriteUINT16(0); err23 != nil {
		return err23
	}
	if err24 := self.binaryStreamWriter.WriteUINT16(0); err24 != nil {
		return err24
	}

	return nil
}

const PACKET_MESSAGE_MAGIC = 2

func (self *PacketWriter) WriteMessageHeader(msgHeader *PacketMessageHeader) (error) {

	/* Write packet message version (2 byte) */
	if err1 := self.binaryStreamWriter.WriteUINT16(PACKET_MESSAGE_MAGIC); err1 != nil {
		return err1
	}

	/* Write origination node (2 byte) */
	if err2 := self.binaryStreamWriter.WriteUINT16(msgHeader.OrigAddr.Node); err2 != nil {
		return err2
	}

	if err3 := self.binaryStreamWriter.WriteUINT16(msgHeader.DestAddr.Node); err3 != nil {
		return err3
	}

	if err4 := self.binaryStreamWriter.WriteUINT16(msgHeader.OrigAddr.Net); err4 != nil {
		return err4
	}

	if err5 := self.binaryStreamWriter.WriteUINT16(msgHeader.DestAddr.Net); err5 != nil {
		return err5
	}

	if err6 := self.binaryStreamWriter.WriteUINT16(msgHeader.Attributes); err6 != nil {
		return err6
	}

	/* Read unused cost fields (2 bytes) */
	if err7 := self.binaryStreamWriter.WriteUINT8(0); err7 != nil {
		return err7
	}
	if err8 := self.binaryStreamWriter.WriteUINT8(0); err8 != nil {
		return err8
	}

	/* Read datetime */
	var pktDateTime []byte = []byte("03 Jan 20  23:51:10\x00")
	if err9 := self.binaryStreamWriter.WriteBytes(pktDateTime); err9 != nil {
		return err9
	}

	/* Read "To" (var bytes) */
	if err10 := self.binaryStreamWriter.WriteZString([]byte(msgHeader.ToUserName)); err10 != nil {
		return err10
	}

	/* Read "From" (var bytes) */
	if err11 := self.binaryStreamWriter.WriteZString([]byte(msgHeader.FromUserName)); err11 != nil {
		return err11
	}

	/* Read "Subject" */
	if err12 := self.binaryStreamWriter.WriteZString([]byte(msgHeader.Subject)); err12 != nil {
		return err12
	}

	return nil
}

func (self *PacketWriter) WriteMessage(msgBody *MessageBody) (error) {

	/* Step 1. Write area */
	var areaName string = msgBody.GetArea()
	if areaName != "" {
		self.binaryStreamWriter.WriteBytes([]byte("AREA"))
		self.binaryStreamWriter.WriteBytes([]byte(":"))
		self.binaryStreamWriter.WriteBytes([]byte(areaName))
		self.binaryStreamWriter.WriteBytes([]byte("\x0D"))
	}

	/* Step 2. Write kludges */
	for _, k := range msgBody.Kludges {
		self.binaryStreamWriter.WriteBytes([]byte("\x01"))
		self.binaryStreamWriter.WriteBytes([]byte(k.Name))
		self.binaryStreamWriter.WriteBytes([]byte(":"))
		self.binaryStreamWriter.WriteBytes([]byte(k.Value))
		self.binaryStreamWriter.WriteBytes([]byte("\x0D"))
	}

	/* Step 3. Write message body */
	if err1 := self.binaryStreamWriter.WriteZString([]byte(msgBody.Body)); err1 != nil {
		return err1
	}

	return nil
}

func (self *PacketWriter) Close() {
	/* Close binary writer */
	self.binaryStreamWriter.Close()
	self.binaryStreamWriter = nil
	/**/
	if err1 := self.streamWriter.Flush(); err1 != nil {
		panic(err1)
	}
	self.streamWriter = nil
	/* Close OS stream */
	self.stream.Close()
	self.stream = nil
}
