package packet

import "log"

type PacketWriterExtension struct {
}

func NewPacketWriterExtension() *PacketWriterExtension {
	return new(PacketWriterExtension)
}


func (self PacketWriterExtension) WritePacketHeaderFill(writer *BinaryWriter, header *PacketHeader) error {

	log.Printf("Write extension section")

	fill := make([]byte, 20)
	if err13 := writer.WriteBytes(fill); err13 != nil {
		return err13
	}

	return nil

}
