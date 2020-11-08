package packet

type PacketWriterExtension struct {
}

func NewPacketWriterExtension() *PacketWriterExtension {
	return new(PacketWriterExtension)
}

/// WritePacketHeaderExtension write packet 20 byte extension section
func (self PacketWriterExtension) WritePacketHeaderExtension(writer *BinaryWriter, pktHeader *PacketHeader) error {

	/* FSC-0039 - Filler - 4 Byte */
	mem := make([]byte, 4)
	err1 := writer.WriteBytes(mem)
	if err1 != nil {
		return err1
	}

	/* FSC-0039 - PrdCodH - 1 Byte */
	err2 := writer.WriteUINT8(0)
	if err2 != nil {
		return err2
	}

	/* FSC-0039 - PVMinor - 1 Byte */
	err3 := writer.WriteUINT8(0)
	if err3 != nil {
		return err3
	}

	/* FSC-0039 - CapWord - 2 Byte */
	var capWord uint16 = 1
	err4 := writer.WriteUINT16(capWord)
	if err4 != nil {
		return err4
	}

	/* FSC-0039 - OrigZone - 2 Int */
	err5 := writer.WriteUINT16(pktHeader.OrigZone)
	if err5 != nil {
		return err5
	}

	/* FSC-0039 - DestZone - 2 Int */
	err6 := writer.WriteUINT16(pktHeader.DestZone)
	if err6 != nil {
		return err6
	}

	/* FSC-0039 - OrigPoint - 2 Int */
	err7 := writer.WriteUINT16(pktHeader.OrigPoint)
	if err7 != nil {
		return err7
	}

	/* FSC-0039 - DestPoint - 2 Int */
	err8 := writer.WriteUINT16(pktHeader.DestPoint)
	if err8 != nil {
		return err8
	}

	/* FSC-0039 - ProdData - 4 Long */
	var prodData uint32 = 0
	err9 := writer.WriteUINT32(prodData)
	if err9 != nil {
		return err9
	}

	return nil

}
