package packet

type PacketWriterExtension struct {
}

func NewPacketWriterExtension() *PacketWriterExtension {
	return new(PacketWriterExtension)
}

/// WritePacketHeaderExtension write packet 20 byte extension section
func (self PacketWriterExtension) WritePacketHeaderExtension(writer *BinaryWriter, pktHeader *PacketHeader) error {

	/* FSC-0048 - AuxNet - 2 Byte */
	var auxNet uint16 = pktHeader.OrigNet
	err1 := writer.WriteUINT16(auxNet)
	if err1 != nil {
		return err1
	}

	/* FSC-0048 - CWvalidationCopy - 2 Byte */
	var CWvalidationCopy uint16 = 0x100
	err2 := writer.WriteUINT16(CWvalidationCopy)
	if err2 != nil {
		return err2
	}

	/* FSC-0048 - ProductCode - 1 Byte */
	var productCode uint8 = 0
	err3 := writer.WriteUINT8(productCode)
	if err3 != nil {
		return err3
	}

	/* FSC-0048 - Revision - 1 Byte */
	var revision uint8 = 0
	err4 := writer.WriteUINT8(revision)
	if err4 != nil {
		return err4
	}

	/* FSC-0048 - CapabilWord - 2 Byte */
	var capabilWord uint16 = 1
	err5 := writer.WriteUINT16(capabilWord)
	if err5 != nil {
		return err5
	}

	/* FSC-0048 - OrigZone - 2 Byte */
	err6 := writer.WriteUINT16(pktHeader.OrigZone)
	if err6 != nil {
		return err6
	}

	/* FSC-0048 - DestZone - 2 Byte */
	err7 := writer.WriteUINT16(pktHeader.DestZone)
	if err7 != nil {
		return err7
	}

	/* FSC-0048 - OrigPoint - 2 Byte */
	err8 := writer.WriteUINT16(pktHeader.OrigPoint)
	if err8 != nil {
		return err8
	}

	/* FSC-0048 - DestPoint - 2 Byte */
	err9 := writer.WriteUINT16(pktHeader.DestPoint)
	if err9 != nil {
		return err9
	}

	/* FSC-0048 - Product Specific Data - 4 Bytes */
	var productData uint32 = 0
	err10 := writer.WriteUINT32(productData)
	if err10 != nil {
		return err10
	}

	return nil

}
