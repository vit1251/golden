package packet

import (
	"encoding/binary"
	"io"
)

type PacketWriterNew struct {
	stream             io.Writer
	extension          interface{}
}

func NewPacketWriterNew(stream io.Writer) (*PacketWriterNew, error) {

	newPacketWriterNew := new(PacketWriterNew)
	newPacketWriterNew.stream = stream
	newPacketWriterNew.extension = nil

	return newPacketWriterNew, nil
}

func (self PacketWriterNew) WritePacketHeader(pktHeader *PacketHeader) error {

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

	/* Populate */

	packetHeader := pktHeaderNew{}
	packetHeader.OrigNode = pktHeader.OrigNode
	packetHeader.DestNode = pktHeader.DestZone

	err1 := binary.Write(self.stream, binary.LittleEndian, &packetHeader)
	if err1 != nil {
		return err1
	}

	/* Populate extension */
	if self.extension == nil {
		type pktHeaderNewExtension struct {
			Fill [20]byte
		}
		packetHeaderExtension := pktHeaderNewExtension{}
		err2 := binary.Write(self.stream, binary.LittleEndian, &packetHeaderExtension)
		if err2 != nil {
			return err2
		}
	} else {
		panic("not implement")
	}

	return nil

}