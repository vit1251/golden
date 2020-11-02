package packet

import (
	"errors"
	"github.com/vit1251/golden/pkg/fidotime"
	"io"
	"log"
)

type PacketReader struct {
	sourceStream io.Reader
	binaryStreamReader *BinaryReader
	extension *PacketReaderExtension
}

func NewPacketReader(stream io.Reader) *PacketReader {

	/* Create packet reader */
	pr := new(PacketReader)
	pr.sourceStream = stream
	pr.extension = NewPacketReaderExtension()

	/* Create binary stream reader */
	binaryStreamReader := NewBinaryReader(stream)
	pr.binaryStreamReader = binaryStreamReader

	return pr
}

func (self *PacketReader) ReadPacketHeader() (*PacketHeader, error) {

	/* Create packet header */
	pktHeader := new(PacketHeader)

	/* Read orginator node address (2 byte) */
	if value, err1 := self.binaryStreamReader.ReadUINT16(); err1 != nil {
		return nil, err1
	} else {
		pktHeader.OrigNode = value
	}

	/* Read destination node address (2 byte) */
	if value, err2 := self.binaryStreamReader.ReadUINT16(); err2 != nil {
		return nil, err2
	} else {
		pktHeader.DestNode = value
	}

	if value, err1 := self.binaryStreamReader.ReadUINT16(); err1 != nil {
		return nil, err1
	} else {
		pktHeader.Year = value
	}
	if value, err2 := self.binaryStreamReader.ReadUINT16(); err2 != nil {
		return nil, err2
	} else {
		pktHeader.Month = value + 1
	}
	if value, err3 := self.binaryStreamReader.ReadUINT16(); err3 != nil {
		return nil, err3
	} else {
		pktHeader.Day = value
	}
	if value, err4 := self.binaryStreamReader.ReadUINT16(); err4 != nil {
		return nil, err4
	} else {
		pktHeader.Hour = value
	}
	if value, err5 := self.binaryStreamReader.ReadUINT16(); err5 != nil {
		return nil, err5
	} else {
		pktHeader.Minute = value
	}
	if value, err6 := self.binaryStreamReader.ReadUINT16(); err6 != nil {
		return nil, err6
	} else {
		pktHeader.Second = value
	}

	/* Read packet baud */
	if _, err4 := self.binaryStreamReader.ReadUINT16(); err4 != nil {
		return nil, err4
	} else {
		// ignore
	}

	/* Read packet version (2 byte) */
	if pktVersion, err5 := self.binaryStreamReader.ReadUINT16(); err5 != nil {
		return nil, err5
	} else {
		if pktVersion != 2 {
			return nil, errors.New("invalid packet version")
		}
	}

	/* Origination network (2 byte) */
	if value, err6 := self.binaryStreamReader.ReadUINT16(); err6 != nil {
		return nil, err6
	} else {
		pktHeader.OrigNet = value
	}

	/* Destination network (2 byte) */
	if value, err7 := self.binaryStreamReader.ReadUINT16(); err7 != nil {
		return nil, err7
	} else {
		pktHeader.DestNet = value
	}

	/* Read product version (2 byte)*/
	if _, err1 := self.binaryStreamReader.ReadUINT8(); err1 != nil {
		return nil, err1
	}

	if _, err2 := self.binaryStreamReader.ReadUINT8(); err2 != nil {
		return nil, err2
	}

	/* Read packet password (8 byte) */
	if value, err10 := self.binaryStreamReader.ReadBytes(8); err10 != nil {
		return nil, err10
	} else {
		pktHeader.PktPassword = value
	}

	/* Read packet zone (2 byte) */
	if value, err11 := self.binaryStreamReader.ReadUINT16(); err11 != nil {
		return nil, err11
	} else {
		pktHeader.OrigZone = value
	}

	/* Read packet zone (2 byte) */
	if value, err12 := self.binaryStreamReader.ReadUINT16(); err12 != nil {
		return nil, err12
	} else {
		pktHeader.DestZone = value
	}

	/* Read packet fill (20 byte) */
	if self.extension != nil {
		if err1 := self.extension.ReadPacketHeaderFill(self.binaryStreamReader, pktHeader); err1 != nil {
			return nil, err1
		}
	} else {
		if _, err13 := self.binaryStreamReader.ReadBytes(20); err13 != nil {
			return nil, err13
		}
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
