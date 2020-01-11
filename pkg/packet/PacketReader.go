package packet

import (
	"io"
	"os"
	"log"
	"bufio"
	"errors"
)

type PacketReader struct {
	stream                   *os.File          /* .. */
	streamReader             *bufio.Reader     /* .. */
	binaryStreamReader       *BinaryReader     /* .. */
}

func NewPacketReader(name string) (*PacketReader, error) {

	/* Create packet reader */
	pr := new(PacketReader)

	/* Create native OS stream */
	if stream, err := os.Open(name); err != nil {
		return nil, err
	} else {
		pr.stream = stream
	}

	/* Create cache stream */
	pr.streamReader = bufio.NewReader(pr.stream)

	/* Create binary stream reader */
	if binaryStreamReader, err := NewBinaryReader(pr.streamReader); err != nil {
		pr.Close()
		return nil, err
	} else {
		pr.binaryStreamReader = binaryStreamReader
	}

	/* Done */
	return pr, nil
}

func (self *PacketReader) readPacketHeaderDateTime() (*PktDateTime, error) {

	/* New packet date time */
	pktDateTime := new(PktDateTime)

	/* Reading */
	if value, err1 := self.binaryStreamReader.ReadUINT16(); err1 != nil {
		return nil, err1
	} else {
		pktDateTime.Year = value
	}
	if value, err2 := self.binaryStreamReader.ReadUINT16(); err2 != nil {
		return nil, err2
	} else {
		pktDateTime.Mon = value
	}
	if value, err3 := self.binaryStreamReader.ReadUINT16(); err3 != nil {
		return nil, err3
	} else {
		pktDateTime.MDay = value
	}
	if value, err4 := self.binaryStreamReader.ReadUINT16(); err4 != nil {
		return nil, err4
	} else {
		pktDateTime.Hour = value
	}
	if value, err5 := self.binaryStreamReader.ReadUINT16(); err5 != nil {
		return nil, err5
	} else {
		pktDateTime.Min = value
	}
	if value, err6 := self.binaryStreamReader.ReadUINT16(); err6 != nil {
		return nil, err6
	} else {
		pktDateTime.Sec = value
	}

	return pktDateTime, nil
}

func (self *PacketReader) readPacketHeaderCapatiblityBytes() (uint8, uint8, error) {

	var capByte1 uint8
	var capByte2 uint8

	if value, err1 := self.binaryStreamReader.ReadUINT8(); err1 != nil {
		return capByte1, capByte2, err1
	} else {
		capByte1 = value
	}

	if value, err2 := self.binaryStreamReader.ReadUINT8(); err2 != nil {
		return capByte1, capByte2, err2
	} else {
		capByte2 = value
	}

	return capByte1, capByte2, nil
}

func (self *PacketReader) readPacketHeaderProductVersion() (error) {

	if _, err1 := self.binaryStreamReader.ReadUINT8(); err1 != nil {
		return err1
//	} else {
//		header.hiProductCode = value
	}

	if _, err2 := self.binaryStreamReader.ReadUINT8(); err2 != nil {
		return err2
//	} else {
//		header.minorProductRev = value
	}

	return nil
}

func (self *PacketReader) ReadPacketHeader() (*PacketHeader, error) {

	/* Create packet header */
	pktHeader := new(PacketHeader)

	/* Read orginator node address (2 byte) */
	if value, err1 := self.binaryStreamReader.ReadUINT16(); err1 != nil {
		return nil, err1
	} else {
		pktHeader.OrigAddr.Node = value
	}

	/* Read destination node address (2 byte) */
	if value, err2 := self.binaryStreamReader.ReadUINT16(); err2 != nil {
		return nil, err2
	} else {
		pktHeader.DestAddr.Node = value
	}

	/* Read packet create (12 byte) */
	if value, err3 := self.readPacketHeaderDateTime(); err3 != nil {
		return nil, err3
	} else {
		pktHeader.pktCreated = *value
	}

	/* Read unused (2 byte) */
	if _, err4 := self.binaryStreamReader.ReadUINT16(); err4 != nil {
		return nil, err4
	}

	/* Read pakcet version (2 byte) */
	if pktVersion, err5 := self.binaryStreamReader.ReadUINT16(); err5 != nil {
		return nil, err5
	} else {
		if pktVersion != 2 {
			return nil, errors.New("Invalid pkt version")
		}
	}

	/* Origination network (2 byte) */
	if value, err6 := self.binaryStreamReader.ReadUINT16(); err6 != nil {
		return nil, err6
	} else {
		pktHeader.OrigAddr.Net = value
	}

	/* Destination network (2 byte) */
	if value, err7 := self.binaryStreamReader.ReadUINT16(); err7 != nil {
		return nil, err7
	} else {
		pktHeader.DestAddr.Net = value
	}

	/* Read product version (2 byte)*/
	if err8 := self.readPacketHeaderProductVersion(); err8 != nil {
		return nil, err8
	}

	/* Read packet password (8 byte) */
	if value, err10 := self.binaryStreamReader.ReadBytes(8); err10 != nil {
		return nil, err10
	} else {
		pktHeader.pktPassword = value
	}

	/* Read packet zone (2 byte) */
	if value, err11 := self.binaryStreamReader.ReadUINT16(); err11 != nil {
		return nil, err11
	} else {
		pktHeader.OrigAddr.Zone = value
	}

	/* Read packet zone (2 byte) */
	if value, err12 := self.binaryStreamReader.ReadUINT16(); err12 != nil {
		return nil, err12
	} else {
		pktHeader.DestAddr.Zone = value
	}

	/* Read auxNet (2 byte) */
	if value, err13 := self.binaryStreamReader.ReadUINT16(); err13 != nil {
		return nil, err13
	} else {
		pktHeader.auxNet = value
	}

	/* Read capatiblity bytes (2 byte) */
	if capByte1, capByte2, err14 := self.readPacketHeaderCapatiblityBytes(); err14 != nil {
		return nil, err14
	} else {
		pktHeader.capatiblityByte1 = capByte1
		pktHeader.capatiblityByte2 = capByte2
	}

	/* Read product version (2 byte) */
	if err15 := self.readPacketHeaderProductVersion(); err15 != nil {
		return nil, err15
	}

	/* Read capability word (2 byte) */
	if value, err16 := self.binaryStreamReader.ReadUINT16(); err16 != nil {
		return nil, err16
	} else {
		pktHeader.capabilityWord = value
	}

	/* Check capatibility */
	if ok := pktHeader.IsCapatiblity(); !ok {
		return nil, errors.New("Packet capatiblity error")
	}

	/* Read additional zone info (2 byte) */
	if _, err19 := self.binaryStreamReader.ReadUINT16(); err19 != nil {
		return nil, err19
	}
	/* Read additional zone info (2 byte) */
	if _, err20 := self.binaryStreamReader.ReadUINT16(); err20 != nil {
		return nil, err20
	}

	/* Read point (2 byte) */
	if value, err21 := self.binaryStreamReader.ReadUINT16(); err21 != nil {
		return nil, err21
	} else {
		pktHeader.OrigAddr.Point = value
	}

	/* Read point (2 byte) */
	if value, err22 := self.binaryStreamReader.ReadUINT16(); err22 != nil {
		return nil, err22
	} else {
		pktHeader.DestAddr.Point = value
	}

	/* Read production data (2 byte) */
	if _, err23 := self.binaryStreamReader.ReadUINT16(); err23 != nil {
		return nil, err23
	}
	if _, err24 := self.binaryStreamReader.ReadUINT16(); err24 != nil {
		return nil, err24
	}

	/* Determine OPUS or FSC */
//	if header.origAddr.net == 0xFFFF {
//		if (header->origAddr.point) {
//			header->origAddr.net = header->auxNet;
//		} else {
//			/*  not in FSC ! */
//			header->origAddr.net = header->destAddr.net;
//		}
//	}

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
			return nil, errors.New("Invalid packet message version")
		}
	}

	/* Read origination node (2 byte) */
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
	} else {
		msgHeader.OrigAddr.Node = value
	}
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
	} else {
		msgHeader.DestAddr.Node = value
	}
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
	} else {
		msgHeader.OrigAddr.Net = value
	}
	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
	} else {
		msgHeader.DestAddr.Net = value
	}

	if value, err := self.binaryStreamReader.ReadUINT16(); err != nil {
		msgHeader.Attributes = value
	}

	/* Read unused cost fields (2bytes) */
	if _, err := self.binaryStreamReader.ReadUINT8(); err != nil {
	} else {
	}
	if _, err := self.binaryStreamReader.ReadUINT8(); err != nil {
	} else {
	}

	/* Read datetime */
	if value, err := self.binaryStreamReader.ReadBytes(20); err != nil {
	} else {

		log.Printf("datetime = %s", value)

		/* Create new one parser */
		parser := NewDateParser()
		if date, err1 := parser.Parse(value); err1 != nil {
		} else {
			if date == nil {
				// TODO - error handling ...
			} else {
				if stamp, err2 := date.Time(); err2 != nil {
				} else {
					msgHeader.Time = stamp
				}
			}
		}

	}
	/* Read "To" (var bytes) */
	if value, err := self.binaryStreamReader.ReadZString(); err != nil {
	} else {
//		log.Printf("ToUserName = %s", value)
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msgHeader.ToUserName = string(content)
		}
	}

	/* Read "From" (var bytes) */
	if value, err := self.binaryStreamReader.ReadZString(); err != nil {
	} else {
//		log.Printf("FromUserName = %s", value)
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msgHeader.FromUserName = string(content)
		}
	}

	/* Read "Subject" */
	if value, err := self.binaryStreamReader.ReadZString(); err != nil {
	} else {
//		log.Printf("Subject = %s", value)
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msgHeader.Subject = string(content)
		}
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

	/**/
	if self.binaryStreamReader != nil {
		self.binaryStreamReader.Close()
		self.binaryStreamReader = nil
	}

	/**/
	if self.streamReader != nil {
//		self.streamReader.Flush()
		self.streamReader = nil
	}

	/**/
	if self.stream != nil {
		self.stream.Close()
		self.stream = nil
	}

}
