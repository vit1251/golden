package packet

import (
	"bufio"
	"io"
	"os"
)

type PacketWriter struct {
	stream                    io.Writer    /* Native OS stream */
	binaryStreamWriter       *BinaryWriter /* ...              */
}

func NewPacketWriter(name string) (*PacketWriter, error) {
	/* Crete new writer */
	writer := new(PacketWriter)
	/* Create native OS stream */
	if stream, err := os.Create(name); err != nil {
		return nil, err
	} else {
		writer.stream = stream
	}
	/* Create cache stream */
	streamWriter := bufio.NewWriter(writer.stream)
	/* Create binary stream reader */
	if binaryStreamWriter, err := NewBinaryWriter(streamWriter); err != nil {
		writer.Close()
		return nil, err
	} else {
		writer.binaryStreamWriter = binaryStreamWriter
	}
	/* Done */
	return writer, nil
}

func (self *PacketWriter) WriteHeader(pktHeader *PacketHeader) {
}

func (self *PacketWriter) WriteMessage(pktMessage *PacketMessage) {
}

func (self *PacketWriter) WriteMessageBody() {
}

func (self *PacketWriter) Close() {
	//self.stream.Close()
}
