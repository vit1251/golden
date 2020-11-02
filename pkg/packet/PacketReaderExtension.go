package packet

import "log"

type PacketReaderExtension struct {
}

func NewPacketReaderExtension() *PacketReaderExtension {
	return new(PacketReaderExtension)
}

func (self PacketReaderExtension) ReadPacketHeaderFill(stream *BinaryReader, pktHeader *PacketHeader) error {

	log.Printf("Read 20 bytes in extension section")

//	_, err0 := stream.ReadBytes(20)
//	if err0 != nil {
//		return err0
//	}

	/* FSC-0039 - Filler - 4 Byte */
	_, err1 := stream.ReadBytes(4)
	if err1 != nil {
		return err1
	}

	/* FSC-0039 - PrdCodH - 1 Byte */
	_, err2 := stream.ReadUINT8()
	if err2 != nil {
		return err2
	}

	/* FSC-0039 - PVMinor - 1 Byte */
	_, err3 := stream.ReadUINT8()
	if err3 != nil {
		return err3
	}

	/* FSC-0039 - CapWord - 2 Byte */
	capWord, err4 := stream.ReadUINT16()
	if err4 != nil {
		return err4
	}
	log.Printf("capWord = %x", capWord)

	/* FSC-0039 - OrigZone - 2 Int */
	origZone, err5 := stream.ReadUINT16()
	if err5 != nil {
		return err5
	}
	pktHeader.OrigZone = origZone

	/* FSC-0039 - DestZone - 2 Int */
	destZone, err6 := stream.ReadUINT16()
	if err6 != nil {
		return err6
	}
	pktHeader.DestZone = destZone

	/* FSC-0039 - OrigPoint - 2 Int */
	origPoint, err7 := stream.ReadUINT16()
	if err7 != nil {
		return err7
	}
	pktHeader.OrigPoint = origPoint

	/* FSC-0039 - DestPoint - 2 Int */
	destPoint, err8 := stream.ReadUINT16()
	if err8 != nil {
		return err8
	}
	pktHeader.DestPoint = destPoint

	/* FSC-0039 - ProdData - 4 Long */
	if _, err4 := stream.ReadUINT32(); err4 != nil {
		return err4
	}

	return nil
}
