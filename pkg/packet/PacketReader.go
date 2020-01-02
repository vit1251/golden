package packet

import (
	"io"
	"os"
	"log"
	"bufio"
	"errors"
	"fmt"
)

func (self *PacketHeader) packetDateTimeRead(stream *BinaryReader) (error) {
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		self.pktCreated.Year = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		self.pktCreated.Mon = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		self.pktCreated.MDay = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		self.pktCreated.Hour = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		self.pktCreated.Min = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		self.pktCreated.Sec = value
	}
	return nil
}

func packetHeaderRead(stream *BinaryReader) (*PacketHeader, error) {

	/* Create packet header */
	header := new(PacketHeader)

	/* Read orginator node address */
	if value, err := stream.ReadUINT16(); err != nil {
		return nil, errors.New("Fail on read originator node address")
	} else {
		log.Printf("Origination node address: %v", value)
		header.OrigAddr.Node = value
	}
	/* Read destination node address */
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		header.DestAddr.Node = value
	}
	/*  12 bytes */
	if err := header.packetDateTimeRead(stream); err != nil {
		return nil, errors.New("Fail on read pktDateTime")
	}
	/* read 2 bytes for the unused baud field */
	if _, err := stream.ReadUINT16(); err != nil {
		return nil, errors.New("Fail on read unused baud field")
	}
	if pktVersion, err := stream.ReadUINT16(); err != nil {
		return nil, errors.New("Fail on read packet version")
	} else {
		if pktVersion != 2 {
			log.Printf("Packet version is %d", pktVersion)
			return nil, errors.New("Invalid pkt version")
		}
	}
	/* Origination network */
	if value, err := stream.ReadUINT16(); err != nil {
		return nil, errors.New("Fail on read origination network")
	} else {
		header.OrigAddr.Net = value
	}
	/* Destination network */
	if value, err := stream.ReadUINT16(); err != nil {
		return nil, errors.New("Fail on read destination network")
	} else {
		header.DestAddr.Net = value
	}
	if value, err := stream.ReadByte(); err != nil {
	} else {
		header.loProductCode = value
	}
	if value, err := stream.ReadByte(); err != nil {
	} else {
		header.majorProductRev = value
	}
	if value, err := stream.ReadString(8); err != nil {
	} else {
		log.Printf("pktPassword = %s", value)
//		readPktPassword(pkt, (UCHAR *)header->pktPassword); /*  8 bytes */
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		header.OrigAddr.Zone = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		header.DestAddr.Zone = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		log.Printf("auxNet = %d", value)
//		header.auxNet = getUINT16(pkt);
	}
	if value, err := stream.ReadByte(); err != nil {
	} else {
		header.capatiblityByte1 = value
	}
	if value, err := stream.ReadByte(); err != nil {
	} else {
		header.capatiblityByte2 = value
	}
	if value, err := stream.ReadByte(); err != nil {
	} else {
		header.hiProductCode = value
	}
	if value, err := stream.ReadByte(); err != nil {
	} else {
		header.minorProductRev = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		header.capabilityWord = value
	}
	/* Check capatibility */
	var capWord uint16 = uint16(header.capatiblityByte1 << 8) + uint16(header.capatiblityByte2)
	if capWord != header.capabilityWord {
		return nil, errors.New("CapabilityWord error in following pkt!")
	}
	/* Read additional zone info */
	stream.ReadUINT16()
	stream.ReadUINT16()
	/* Read point */
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		header.OrigAddr.Point = value
	}
	/* Read point */
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		header.DestAddr.Point = value
	}
	/* Read production data */
	stream.ReadUINT16()
	stream.ReadUINT16()
	/* Determine OPUS or FSC */
//	if header.origAddr.net == 0xFFFF {
//		if (header->origAddr.point) {
//			header->origAddr.net = header->auxNet;
//		} else {
//			/*  not in FSC ! */
//			header->origAddr.net = header->destAddr.net;
//		}
//	}
//	if (header->origAddr.zone == 0) {
//		for (capWord=0; capWord<config->addrCount; capWord++) {
//			if (header->origAddr.net==config->addr[capWord].net) {
//				header->origAddr.zone = config->addr[capWord].zone;
//				break;
//			}
//		}
//		if (header->origAddr.zone==0) header->origAddr.zone=config->addr[0].zone;
//	}
//	if (header->destAddr.zone == 0) {
//		for (capWord=0; capWord<config->addrCount; capWord++) {
//			if (header->destAddr.net==config->addr[capWord].net) {
//				header->destAddr.zone = config->addr[capWord].zone;
//				break;
//			}
//		}
//		if (header->destAddr.zone==0) header->destAddr.zone=config->addr[0].zone;
//	}

	return header, nil
}

func packetMessageRead(stream *BinaryReader) (*PacketMessage, error) {

	msg := new(PacketMessage)

	/* Check version */
	if value, err := stream.ReadUINT16(); err != nil {
		return nil, err
	} else {
		if value == 0 {
			return nil, io.EOF
		} else if value == 2 {
			/* Valid */
		} else {
			return nil, fmt.Errorf("Invalid packet message header: version = %d (offset = %d)", value, stream.Offset())
		}
	}

	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		msg.OrigAddr.Node = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		msg.DestAddr.Node = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		msg.OrigAddr.Net = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
	} else {
		msg.DestAddr.Net = value
	}
	if value, err := stream.ReadUINT16(); err != nil {
		msg.Attributes = value
	}
	/* Read unused cost fields (2bytes) */
	if _, err := stream.ReadByte(); err != nil {
	} else {
	}
	if _, err := stream.ReadByte(); err != nil {
	} else {}
	/* Read datetime */
	if value, err := stream.ReadString(20); err != nil {
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
					msg.Time = stamp
				}
			}
		}

	}
	/* Read To header */
	if value, err := stream.ReadUntil('\x00'); err != nil {
	} else {
//		log.Printf("ToUserName = %s", value)
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msg.ToUserName = string(content)
		}
	}
	/* Read From header */
	if value, err := stream.ReadUntil('\x00'); err != nil {
	} else {
//		log.Printf("FromUserName = %s", value)
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msg.FromUserName = string(content)
		}
	}
	/* Read Subject */
	if value, err := stream.ReadUntil('\x00'); err != nil {
	} else {
//		log.Printf("Subject = %s", value)
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msg.Subject = string(content)
		}
	}
	/* Read Text */
	if value, err := stream.ReadUntil('\x00'); err != nil {
	} else {
		//
		msg.RAW = value
		//
		if content, err1 := DecodeText(value); err1 != nil {
		} else {
			msg.Text = string(content)

		}
	}

	return msg, nil
}

type PacketReader struct {
	stream                   io.Reader
	binaryStreamReader       *BinaryReader
	packetHeader             *PacketHeader
}

func NewPacketReader(name string) (*PacketReader, error) {
	/* Create packet reader */
	reader := new(PacketReader)
	/* Create native OS stream */
	if stream, err := os.Open(name); err != nil {
		return nil, err
	} else {
		reader.stream = stream
	}
	/* Create cache stream */
	streamReader := bufio.NewReader(reader.stream)
	/* Create binary stream reader */
	if binaryStreamReader, err := NewBinaryReader(streamReader); err != nil {
		reader.Close()
		return nil, err
	} else {
		reader.binaryStreamReader = binaryStreamReader
	}
	/* Read packet header */
	if packetHeader, err := packetHeaderRead(reader.binaryStreamReader); err != nil {
		reader.Close()
		return nil, err
	} else {
		reader.packetHeader = packetHeader
	}
	log.Printf("Process packet %s: source = %s destination = %s", name, reader.packetHeader.OrigAddr.String(), reader.packetHeader.DestAddr.String())
	/* Done */
	return reader, nil
}

type FidoMessage struct {
	Headers map[string]string
	Text    []byte
}

func (self *PacketReader) Scan() <-chan Message {
	c := make(chan Message)
	go func() {
		for {

			/* Parse message */
			msg, err3 := packetMessageRead(self.binaryStreamReader)
			if err3 != nil {
				if err3 == io.EOF {
					break
				} else {
					log.Fatal(err3)
					break
				}
			}

			/* Parse message content */
			mbp, err4 := NewMessageBodyParser()
			if err4 != nil {
				log.Fatal(err4)
				break
			}

			mb, err5 := mbp.Parse(msg.RAW)
			if err5 != nil {
				log.Fatal(err4)
				break
			}

			log.Printf("mb = %q", mb)

			/* Area name */
			var areaName string = "BAD"
			if area, ok := mb.Kludges["AREA"]; ok {
				areaName = area
			}

			/* Create message on storage */
			outMsg := NewMessage()
			outMsg.From = msg.FromUserName
			outMsg.To = msg.ToUserName
			outMsg.Subject = msg.Subject
			outMsg.Area = areaName
			outMsg.Content = mb.Text()
			outMsg.SetTime(msg.Time)
			c <- *outMsg

//			log.Printf("msg = %V", msg)

		}
		close(c)
	}()
	return c
}

func (self *PacketReader) Close() {
//	self.stream.Close()
}
