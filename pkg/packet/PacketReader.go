package packet

import (
	"errors"
	"fmt"
	"github.com/vit1251/golden/pkg/fidotime"
	"io"
	"log"
)

type PacketReader struct {
	sourceStream io.Reader
	binaryStreamReader *BinaryReader
}

func NewPacketReader(stream io.Reader) *PacketReader {

	/* Create packet reader */
	pr := new(PacketReader)
	pr.sourceStream = stream

	/* Create binary stream reader */
	binaryStreamReader := NewBinaryReader(stream)
	pr.binaryStreamReader = binaryStreamReader

	return pr
}

func (self *PacketReader) ReadPacketHeader() (*PacketHeader, error) {

	reader := self.binaryStreamReader

	/* Create packet header */
	pktHeader := new(PacketHeader)

	/* Read orig node address (2 byte) */
	fromnode, err1 := reader.ReadUINT16()
	if err1 != nil {
		return nil, err1
	}
	pktHeader.OrigNode = fromnode

	/* Read dest node address (2 byte) */
	tonode, err2 := reader.ReadUINT16()
	if err2 != nil {
		return nil, err2
	}
	pktHeader.DestNode = tonode

	/* */
	year, err3 := reader.ReadUINT16()
	if err3 != nil {
		return nil, err3
	}
	pktHeader.Year = year

	month, err4 := reader.ReadUINT16()
	if err4 != nil {
		return nil, err4
	}
	pktHeader.Month = month + 1

	day, err5 := reader.ReadUINT16()
	if err5 != nil {
		return nil, err5
	}
	pktHeader.Day = day

	hour, err6 := reader.ReadUINT16()
	if err6 != nil {
		return nil, err6
	}
	pktHeader.Hour = hour

	minute, err7 := reader.ReadUINT16()
	if err7 != nil {
		return nil, err7
	}
	pktHeader.Minute = minute

	second, err8 := reader.ReadUINT16()
	if err8 != nil {
		return nil, err8
	}
	pktHeader.Second = second

	/* Read packet baud */
	_, err9 := reader.ReadUINT16()
	if err9 != nil {
		return nil, err9
	}

	/* Read packet version (2 byte) */
	pktVersion, err10 := reader.ReadUINT16()
	if err10 != nil {
		return nil, err10
	}
	if pktVersion != 2 {
		return nil, errors.New("invalid packet version")
	}

	/* Origination network (2 byte) */
	fromnet, err11 := reader.ReadUINT16()
	if err11 != nil {
		return nil, err11
	}
	pktHeader.OrigNet = fromnet

	/* Destination network (2 byte) */
	tonet, err12 := reader.ReadUINT16()
	if err12 != nil {
		return nil, err12
	}
	pktHeader.DestNet = tonet

	/* Read product version (2 byte)*/
	_, err13 := reader.ReadUINT8()
	if err13 != nil {
		return nil, err13
	}

	_, err14 := reader.ReadUINT8()
	if err14 != nil {
		return nil, err14
	}

	/* Read packet password (8 byte) */
	password, err15 := reader.ReadBytes(8)
	if err15 != nil {
		return nil, err15
	}
	pktHeader.PktPassword = password

	/* Read packet zone (2 byte) */
	fromzone, err16 := reader.ReadUINT16()
	if err16 != nil {
		return nil, err11
	}
	pktHeader.OrigZone = fromzone

	/* Read packet zone (2 byte) */
	tozone, err17 := reader.ReadUINT16()
	if err17 != nil {
		return nil, err17
	}
	pktHeader.DestZone = tozone

	/* FSC-0048 - AuxNet - 2 Byte */
	auxNet, err18 := reader.ReadUINT16()
	if err18 != nil {
		return nil, err18
	}

	/* FSC-0048 - CWvalidationCopy - 2 Byte */
	CWvalidationCopy, err19 := reader.ReadUINT16()
	if err19 != nil {
		return nil, err19
	}

	/* FSC-0048 - ProductCode - 1 Byte */
	_, err20 := reader.ReadUINT8()
	if err20 != nil {
		return nil, err20
	}

	/* FSC-0048 - Revision - 1 Byte */
	_, err21 := reader.ReadUINT8()
	if err21 != nil {
		return nil, err21
	}

	/* FSC-0048 - CapabilWord - 2 Byte */
	capabilWord, err22 := reader.ReadUINT16()
	if err22 != nil {
		return nil, err22
	}

	/* FSC-0048 - OrigZone - 2 Byte */
	origZone, err23 := reader.ReadUINT16()
	if err23 != nil {
		return nil, err23
	}

	/* FSC-0048 - DestZone - 2 Byte */
	destZone, err24 := reader.ReadUINT16()
	if err24 != nil {
		return nil, err24
	}

	/* FSC-0048 - OrigPoint - 2 Byte */
	origPoint, err25 := reader.ReadUINT16()
	if err25 != nil {
		return nil, err25
	}

	/* FSC-0048 - DestPoint - 2 Byte */
	destPoint, err26 := reader.ReadUINT16()
	if err26 != nil {
		return nil, err26
	}

	/* FSC-0048 - Product Specific Data - 4 Bytes */
	_, err27 := reader.ReadUINT32()
	if err27 != nil {
		return nil, err27
	}

	/* Checks */

	var newCapabilWord uint16 = ((CWvalidationCopy >> 8) + (CWvalidationCopy << 8)) & 0xFFFF
	if (capabilWord != 0) && (capabilWord == newCapabilWord) {
		log.Printf("PacketReaderExtension: Packet process as FSC-0048 packet")
	} else {
		return nil, fmt.Errorf("error in FSC-0048 capatibility word value")
	}

	if pktHeader.OrigZone == 0 {
		pktHeader.OrigZone = origZone
	}
	if pktHeader.DestZone == 0 {
		pktHeader.DestZone = destZone
	}
	pktHeader.OrigPoint = origPoint
	pktHeader.DestPoint = destPoint
	if pktHeader.OrigNet == 65535 {
		pktHeader.OrigNet = auxNet
	}

	return pktHeader, nil
}

func (self *PacketReader) ReadMessageHeader() (*PacketMessageHeader, error) {

	msgHeader := new(PacketMessageHeader)

	/* Read packet message version (2 byte) */
	if value, err1 := self.binaryStreamReader.ReadUINT16(); err1 != nil {
		return nil, err1
	} else {
		if value == 0 {
			return nil, io.EOF
		} else if value == 2 {
			/* Valid */
		} else {
			return nil, errors.New("invalid packet message version")
		}
	}

	/* Read origination node (2 byte) */
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
		return nil, err
	} else {
		msgHeader.OrigAddr.Node = value
	}
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
		return nil, err
	} else {
		msgHeader.DestAddr.Node = value
	}
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
		return nil, err
	} else {
		msgHeader.OrigAddr.Net = value
	}
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
		return nil, err
	} else {
		msgHeader.DestAddr.Net = value
	}

	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
		return nil, err
	} else {
		msgHeader.Attributes = value
	}

	/* Read unused cost fields (2bytes) */
	if _, err := self.binaryStreamReader.ReadUINT8(); err != nil {
		return nil, err
	} else {
	}
	if _, err := self.binaryStreamReader.ReadUINT8(); err != nil {
		return nil, err
	} else {
	}

	/* Read datetime */
	if value, err := self.binaryStreamReader.ReadBytes(20); err != nil {
		return nil, err
	} else {

		log.Printf("datetime = %s", value)

		/* Create new one parser */
		parser := fidotime.NewDateParser()
		if stamp, err1 := parser.Parse(value); err1 != nil {
			return nil, err1
		} else {
			msgHeader.Time = stamp
		}

	}
	/* Read "To" (var bytes) */
	if value, err := self.binaryStreamReader.ReadZString(); err != nil {
		return nil, err
	} else {
		msgHeader.ToUserName = value
	}

	/* Read "From" (var bytes) */
	if value, err := self.binaryStreamReader.ReadZString(); err != nil {
		return nil, err
	} else {
		msgHeader.FromUserName = value
	}

	/* Read "Subject" */
	if value, err := self.binaryStreamReader.ReadZString(); err != nil {
		return nil, err
	} else {
		msgHeader.Subject = value
	}

	return msgHeader, nil
}

func (self *PacketReader) ReadMessage() ([]byte, error) {

	/* Read message body (var bytes) */
	body, err1 := self.binaryStreamReader.ReadZString()
	if err1 != nil {
		return nil, err1
	}

	/* Done */
	return body, nil
}

func (self *PacketReader) Close() {
}
