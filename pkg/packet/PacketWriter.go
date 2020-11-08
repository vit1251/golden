package packet

import (
	"github.com/vit1251/golden/pkg/fidotime"
	"io"
)

type PacketWriter struct {
	streamWriter             io.Writer
	binaryStreamWriter       *BinaryWriter
	extension                *PacketWriterExtension
}

func NewPacketWriter(stream io.Writer) (*PacketWriter, error) {

	/* Crete new packet writer */
	pw := new(PacketWriter)
	pw.streamWriter = stream
	pw.extension = NewPacketWriterExtension()

	/* Create binary stream reader */
	if binaryStreamWriter, err := NewBinaryWriter(stream); err == nil {
		pw.binaryStreamWriter = binaryStreamWriter
	} else {
		return nil, err
	}

	/* Done */
	return pw, nil
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

func (self *PacketWriter) WritePacketHeader(pktHeader *PacketHeader) error {

	/* Write origin node address */
	if err1 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigNode); err1 != nil {
		return err1
	}

	/* Write dest node address */
	if err2 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestNode); err2 != nil {
		return err2
	}

	/* Write packet create (12 byte) */
	if err1 := self.binaryStreamWriter.WriteUINT16(pktHeader.Year); err1 != nil {
		return err1
	}
	if err2 := self.binaryStreamWriter.WriteUINT16(pktHeader.Month); err2 != nil {
		return err2
	}
	if err3 := self.binaryStreamWriter.WriteUINT16(pktHeader.Day); err3 != nil {
		return err3
	}
	if err4 := self.binaryStreamWriter.WriteUINT16(pktHeader.Hour); err4 != nil {
		return err4
	}
	if err5 := self.binaryStreamWriter.WriteUINT16(pktHeader.Minute); err5 != nil {
		return err5
	}
	if err6 := self.binaryStreamWriter.WriteUINT16(pktHeader.Second); err6 != nil {
		return err6
	}

	/* Write baud (2 byte) */
	if err4 := self.binaryStreamWriter.WriteUINT16(0); err4 != nil {
		return err4
	}

	/* Write packet version (2 byte) */
	if err5 := self.binaryStreamWriter.WriteUINT16(PKT_VERSION); err5 != nil {
		return err5
	}

	/* Write orig network (2 byte) */
	if err6 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigNet); err6 != nil {
		return err6
	}

	/* Write dest network (2 byte) */
	if err7 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestNet); err7 != nil {
		return err7
	}

	/* Write prodCode version (1 byte)*/
	if err1 := self.binaryStreamWriter.WriteUINT8(0); err1 != nil {
		return err1
	}

	/* Write serialNo version (1 byte)*/
	if err2 := self.binaryStreamWriter.WriteUINT8(0); err2 != nil {
		return err2
	}

	/* Write packet password (8 byte) */
	pktPassword := make([]byte, 8)
	copy(pktPassword, pktHeader.PktPassword)
	if err10 := self.binaryStreamWriter.WriteBytes(pktPassword); err10 != nil {
		return err10
	}

	/* Write orig zone (2 byte) */
	if err11 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigZone); err11 != nil {
		return err11
	}

	/* Write dest zone (2 byte) */
	if err12 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestZone); err12 != nil {
		return err12
	}

	/* Write packet fill (20 byte) */
	if self.extension != nil {
		if err1 := self.extension.WritePacketHeaderExtension(self.binaryStreamWriter, pktHeader); err1 != nil {
			return err1
		}
	} else {
		fill := make([]byte, 20)
		if err13 := self.binaryStreamWriter.WriteBytes(fill); err13 != nil {
			return err13
		}
	}

	return nil
}

const PACKET_MESSAGE_MAGIC = 2

func (self *PacketWriter) WriteMessageHeader(msgHeader *PacketMessageHeader) error {

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
	newFidoDate := fidotime.NewFidoDate()
	newFidoDate.SetNow()
	var pktDateTime []byte = newFidoDate.FTSC()
	if err9_0 := self.binaryStreamWriter.WriteBytes(pktDateTime); err9_0 != nil {
		return err9_0
	}
	if err9_1 := self.binaryStreamWriter.WriteBytes([]byte("\x00")); err9_1 != nil {
		return err9_1
	}

	/* Read "To" (var bytes) */
	if err10 := self.binaryStreamWriter.WriteZString([]byte(msgHeader.ToUserName), 36 - 1); err10 != nil {
		return err10
	}

	/* Read "From" (var bytes) */
	if err11 := self.binaryStreamWriter.WriteZString([]byte(msgHeader.FromUserName), 36 - 1); err11 != nil {
		return err11
	}

	/* Read "Subject" */
	if err12 := self.binaryStreamWriter.WriteZString([]byte(msgHeader.Subject), 72 - 1); err12 != nil {
		return err12
	}

	return nil
}

const (
	SOH = "\x01"
	CR  = "\x0D"
)

func (self *PacketWriter) WriteMessage(msgBody *MessageBody) error {

	/* Step 1. Write area */
	var areaName string = msgBody.GetArea()
	if areaName != "" {
		self.binaryStreamWriter.WriteBytes([]byte("AREA"))
		self.binaryStreamWriter.WriteBytes([]byte(":"))
		self.binaryStreamWriter.WriteBytes([]byte(areaName))
		self.binaryStreamWriter.WriteBytes([]byte(CR))
	}

	/* Step 2. Write kludges */
	for _, k := range msgBody.kludges {
		self.binaryStreamWriter.WriteBytes(k.Raw)
		self.binaryStreamWriter.WriteBytes([]byte(CR))
	}

	/* Step 3. Write message body */
	msgBodyRaw := msgBody.GetRaw()
	if err1 := self.binaryStreamWriter.WriteZString(msgBodyRaw, 0); err1 != nil {
		return err1
	}

	return nil
}

func (self *PacketWriter) WritePacketEnd() error {

	var packetEndMarker uint16 = 0
	if err6 := self.binaryStreamWriter.WriteUINT16(packetEndMarker); err6 != nil {
		return err6
	}

	return nil

}
