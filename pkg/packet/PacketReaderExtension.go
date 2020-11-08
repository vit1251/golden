package packet

import "fmt"

type PacketReaderExtension struct {
}

func NewPacketReaderExtension() *PacketReaderExtension {
	return new(PacketReaderExtension)
}

// ReadPacketHeaderExtension read packet 20 byte extension section
func (self PacketReaderExtension) ReadPacketHeaderExtension(stream *BinaryReader, pktHeader *PacketHeader) error {

	/* FSC-0048 - AuxNet - 2 Byte */
	_, err1 := stream.ReadUINT16()
	if err1 != nil {
		return err1
	}

	/* FSC-0048 - CWvalidationCopy - 2 Byte */
	CWvalidationCopy, err2 := stream.ReadUINT16()
	if err2 != nil {
		return err2
	}

	/* FSC-0048 - ProductCode - 1 Byte */
	_, err3 := stream.ReadUINT8()
	if err3 != nil {
		return err3
	}

	/* FSC-0048 - Revision - 1 Byte */
	_, err4 := stream.ReadUINT8()
	if err4 != nil {
		return err4
	}

	/* FSC-0048 - CapabilWord - 2 Byte */
	capabilWord, err5 := stream.ReadUINT16()
	if err5 != nil {
		return err5
	}

	/* FSC-0048 - OrigZone - 2 Byte */
	origZone, err6 := stream.ReadUINT16()
	if err6 != nil {
		return err6
	}
	pktHeader.OrigZone = origZone

	/* FSC-0048 - DestZone - 2 Byte */
	destZone, err7 := stream.ReadUINT16()
	if err7 != nil {
		return err7
	}
	pktHeader.DestZone = destZone

	/* FSC-0048 - OrigPoint - 2 Byte */
	origPoint, err8 := stream.ReadUINT16()
	if err8 != nil {
		return err8
	}
	pktHeader.OrigPoint = origPoint

	/* FSC-0048 - DestPoint - 2 Byte */
	destPoint, err9 := stream.ReadUINT16()
	if err9 != nil {
		return err9
	}
	pktHeader.DestPoint = destPoint

	/* FSC-0048 - Product Specific Data - 4 Bytes */
	_, err10 := stream.ReadUINT32()
	if err10 != nil {
		return err10
	}

	/* Check */
	var newCapabilWord uint16 = ((CWvalidationCopy >> 8) + (CWvalidationCopy << 8)) & 0xFFFF
	if (capabilWord != 0) && (capabilWord == newCapabilWord) {
		// ignore
	} else {
		return fmt.Errorf("error in FSC-0048 capatibility word value")
	}

	return nil
}
