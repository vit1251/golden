package packet

import (
	"encoding/binary"
	"io"
)

type PacketReaderNew struct {
	reader    io.Reader
	extension interface{}
}

func NewPacketReaderNew(stream io.Reader) *PacketReaderNew {
	newPacketReaderNew := new(PacketReaderNew)
	newPacketReaderNew.reader = stream
	newPacketReaderNew.extension = nil
	return newPacketReaderNew
}

func (self PacketReaderNew) ReadPacketHeader() (*PacketHeader, error) {

	type pktHeaderNew struct {
		OrigNode uint16
		DestNode uint16
		Year uint16
		Month uint16
		Day uint16
		Hour uint16
		Minute uint16
		Second uint16
		Baud uint16
		PacketType uint16
		OrigNet uint16
		DestNet uint16
		ProdCode uint8
		SerialNo uint8
		Password [8]byte
		OrigZone uint16
		DestZone uint16
	}

	packetHeader := pktHeaderNew{}

	err1 := binary.Read(self.reader, binary.LittleEndian, &packetHeader)
	if err1 != nil {
		return nil, err1
	}

	result := PacketHeader{

		/* Origination */
		OrigZone: packetHeader.OrigZone,
		OrigNet: packetHeader.OrigNet,
		OrigNode: packetHeader.OrigNode,
		OrigPoint: 0,

		/* Destination */
		DestZone: packetHeader.DestZone,
		DestNet: packetHeader.DestNet,
		DestNode: packetHeader.DestNode,
		DestPoint: 0,

		/* DateTime */
		Year: packetHeader.Year,
		Month: packetHeader.Month + 1,
		Day: packetHeader.Day,

		Hour: packetHeader.Hour,
		Minute: packetHeader.Minute,
		Second: packetHeader.Second,

		/* Password */
		PktPassword: packetHeader.Password[:],
	}

	/* Process extension */
	if self.extension == nil {
		type PacketHeaderExtension struct {
			Fill [20]byte
		}
		packetHeaderExtension := PacketHeaderExtension{}
		err2 := binary.Read(self.reader, binary.LittleEndian, &packetHeaderExtension)
		if err2 != nil {
			return nil, err1
		}
	} else {
		panic("not implement")
	}

	return &result, nil

}

func (self PacketReaderNew) ReadMessageHeader() (interface{}, error) {
	return nil, nil
}
