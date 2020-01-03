package packet

import (
	"bufio"
	"os"
)

type PacketWriter struct {
	stream                   *os.File      /* Native OS stream */
	streamWriter             *bufio.Writer /* Cache */
	binaryStreamWriter       *BinaryWriter /* ...              */
}

func NewPacketWriter(name string) (*PacketWriter, error) {

	/* Crete new packet writer */
	pw := new(PacketWriter)

	/* Create native OS stream */
	if stream, err := os.Create(name); err != nil {
		return nil, err
	} else {
		pw.stream = stream
	}

	/* Create cache stream */
	pw.streamWriter = bufio.NewWriter(pw.stream)

	/* Create binary stream reader */
	if binaryStreamWriter, err := NewBinaryWriter(pw.streamWriter); err != nil {
		pw.Close()
		return nil, err
	} else {
		pw.binaryStreamWriter = binaryStreamWriter
	}

	/* Done */
	return pw, nil
}

func (self *PacketWriter) writePacketHeaderDateTime(pktDateTime *PktDateTime) (error) {

	/* Writing */
	if err1 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Year); err1 != nil {
		return err1
	}
	if err2 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Mon); err2 != nil {
		return err2
	}
	if err3 := self.binaryStreamWriter.WriteUINT16(pktDateTime.MDay); err3 != nil {
		return err3
	}
	if err4 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Hour); err4 != nil {
		return err4
	}
	if err5 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Min); err5 != nil {
		return err5
	}
	if err6 := self.binaryStreamWriter.WriteUINT16(pktDateTime.Sec); err6 != nil {
		return err6
	}

	return nil
}

const PKT_VERSION = 2

func (self *PacketWriter) WritePacketHeader(pktHeader *PacketHeader) (error) {

	/* Write orginator node address */
	if err1 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigAddr.Node); err1 != nil {
		return err1
	}
	/* Write destination node address */
	if err2 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestAddr.Node); err2 != nil {
		return err2
	}

	/* Write packet create (12 byte) */
	if err3 := self.writePacketHeaderDateTime(&pktHeader.pktCreated); err3 != nil {
		return err3
	}

	/* Write unused (2 byte) */
	if err4 := self.binaryStreamWriter.WriteUINT16(0); err4 != nil {
		return err4
	}

	/* Write packet version (2 byte) */
	if err5 := self.binaryStreamWriter.WriteUINT16(PKT_VERSION); err5 != nil {
		return err5
	}

	/* Origination network (2 byte) */
	if err6 := self.binaryStreamWriter.WriteUINT16(pktHeader.OrigAddr.Net); err6 != nil {
		return err6
	}

	/* Destination network (2 byte) */
	if err7 := self.binaryStreamWriter.WriteUINT16(pktHeader.DestAddr.Net); err7 != nil {
		return err7
	}

	return nil
}

func (self *PacketWriter) WriteMessageHeader(msgHeader *PacketMessageHeader) (error) {
	return nil
}

func (self *PacketWriter) WriteMessage(msgBody *MessageBody) (error) {
	return nil
}

func (self *PacketWriter) Close() {
	/* Close binary writer */
	self.binaryStreamWriter.Close()
	self.binaryStreamWriter = nil
	/**/
	if err1 := self.streamWriter.Flush(); err1 != nil {
		panic(err1)
	}
	self.streamWriter = nil
	/* Close OS stream */
	self.stream.Close()
	self.stream = nil
}
