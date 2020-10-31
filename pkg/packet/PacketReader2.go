package packet

import (
	"encoding/binary"
	"io"
)

type PacketReader2 struct {
	reader io.Reader
}

func NewPacketReader2(stream io.Reader) *PacketReader2 {
	packetReader2 := new(PacketReader2)
	packetReader2.reader = stream
	return packetReader2
}

func (self PacketReader2) ReadPacketHeader() (*PacketHeader, error) {

	type pktHeader struct {
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
		Fill [20]byte
	}

	packetHeader := pktHeader{}

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

	return &result, nil

}

func (self PacketReader2) ReadMessageHeader() (interface{}, error) {
	return nil, nil
}
