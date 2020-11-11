package packet

import (
	"github.com/vit1251/golden/pkg/fidotime"
	"io"
)

type PacketWriter struct {
	streamWriter             io.Writer
	binaryStreamWriter       *BinaryWriter
}

func NewPacketWriter(stream io.Writer) (*PacketWriter, error) {

	/* Crete new packet writer */
	pw := new(PacketWriter)
	pw.streamWriter = stream

	/* Create binary stream reader */
	if binaryStreamWriter, err := NewBinaryWriter(stream); err == nil {
		pw.binaryStreamWriter = binaryStreamWriter
	} else {
		return nil, err
	}

	/* Done */
	return pw, nil
}

func (self *PacketWriter) WritePacketHeader(pktHeader *PacketHeader) error {

	writer := self.binaryStreamWriter

	/* Write origin node address */
	err1 := writer.WriteUINT16(pktHeader.OrigNode)
	if err1 != nil {
		return err1
	}

	/* Write dest node address */
	err2 := writer.WriteUINT16(pktHeader.DestNode)
	if err2 != nil {
		return err2
	}

	/* Write packet create (12 byte) */
	err3 := writer.WriteUINT16(pktHeader.Year)
	if err3 != nil {
		return err3
	}
	err4 := writer.WriteUINT16(pktHeader.Month)
	if err4 != nil {
		return err4
	}
	err5 := writer.WriteUINT16(pktHeader.Day)
	if err5 != nil {
		return err5
	}
	err6 := writer.WriteUINT16(pktHeader.Hour)
	if err6 != nil {
		return err6
	}
	err7 := writer.WriteUINT16(pktHeader.Minute)
	if err7 != nil {
		return err7
	}
	err8 := writer.WriteUINT16(pktHeader.Second)
	if err8 != nil {
		return err8
	}

	/* Write baud (2 byte) */
	err9 := writer.WriteUINT16(0)
	if err9 != nil {
		return err9
	}

	/* Write packet version (2 byte) */
	var pktVersion uint16 = 2
	err10 := writer.WriteUINT16(pktVersion)
	if err10 != nil {
		return err10
	}

	/* Write orig network (2 byte) */
	var origNet uint16 = pktHeader.OrigNet
	if pktHeader.OrigPoint != 0 {
		origNet = 65535
	}
	err11 := writer.WriteUINT16(origNet)
	if err11 != nil {
		return err11
	}

	/* Write dest network (2 byte) */
	err12 := writer.WriteUINT16(pktHeader.DestNet)
	if err12 != nil {
		return err12
	}

	/* Write prodCode version (1 byte)*/
	err13 := writer.WriteUINT8(0)
	if err13 != nil {
		return err13
	}

	/* Write serialNo version (1 byte)*/
	err14 := writer.WriteUINT8(0)
	if err14 != nil {
		return err14
	}

	/* Write packet password (8 byte) */
	pktPassword := make([]byte, 8)
	copy(pktPassword, pktHeader.PktPassword)
	err15 := writer.WriteBytes(pktPassword)
	if err15 != nil {
		return err15
	}

	/* Write orig zone (2 byte) */
	err16 := writer.WriteUINT16(pktHeader.OrigZone)
	if err16 != nil {
		return err16
	}

	/* Write dest zone (2 byte) */
	err17 := writer.WriteUINT16(pktHeader.DestZone)
	if err17 != nil {
		return err17
	}

	/* FSC-0048 - AuxNet - 2 Byte */
	var auxNet uint16 = 0
	if pktHeader.OrigPoint != 0 {
		auxNet = pktHeader.OrigNet
	}
	err18 := writer.WriteUINT16(auxNet)
	if err18 != nil {
		return err18
	}

	/* FSC-0048 - CWvalidationCopy - 2 Byte */
	var CWvalidationCopy uint16 = 0x100
	err19 := writer.WriteUINT16(CWvalidationCopy)
	if err19 != nil {
		return err19
	}

	/* FSC-0048 - ProductCode - 1 Byte */
	var productCode uint8 = 0
	err20 := writer.WriteUINT8(productCode)
	if err20 != nil {
		return err20
	}

	/* FSC-0048 - Revision - 1 Byte */
	var revision uint8 = 0
	err21 := writer.WriteUINT8(revision)
	if err21 != nil {
		return err21
	}

	/* FSC-0048 - CapabilWord - 2 Byte */
	var capabilWord uint16 = 1
	err22 := writer.WriteUINT16(capabilWord)
	if err22 != nil {
		return err22
	}

	/* FSC-0048 - OrigZone - 2 Byte */
	err23 := writer.WriteUINT16(pktHeader.OrigZone)
	if err23 != nil {
		return err23
	}

	/* FSC-0048 - DestZone - 2 Byte */
	err24 := writer.WriteUINT16(pktHeader.DestZone)
	if err24 != nil {
		return err24
	}

	/* FSC-0048 - OrigPoint - 2 Byte */
	err25 := writer.WriteUINT16(pktHeader.OrigPoint)
	if err25 != nil {
		return err25
	}

	/* FSC-0048 - DestPoint - 2 Byte */
	err26 := writer.WriteUINT16(pktHeader.DestPoint)
	if err26 != nil {
		return err26
	}

	/* FSC-0048 - Product Specific Data - 4 Bytes */
	var productData uint32 = 0
	err27 := writer.WriteUINT32(productData)
	if err27 != nil {
		return err27
	}

	return nil
}

const PACKET_MESSAGE_MAGIC = 2

func (self *PacketWriter) WriteMessageHeader(msgHeader *PackedMessage) error {

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
	if err7 := self.binaryStreamWriter.WriteUINT16(0); err7 != nil {
		return err7
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
