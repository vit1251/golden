package packet

import (
	"github.com/vit1251/golden/pkg/fidotime"
	"io"
)

type PacketWriter struct {
	streamWriter       io.Writer
	binaryStreamWriter *BinaryWriter
}

func NewPacketWriter(stream io.Writer) (*PacketWriter, error) {

	/* Crete new packet writer */
	pw := new(PacketWriter)
	pw.streamWriter = stream

	/* Create binary stream reader */
	if binaryStreamWriter, err := NewBinaryWriter(stream); err == nil {
		pw.binaryStreamWriter = binaryStreamWriter
	} else {
		return nil, err
	}

	/* Done */
	return pw, nil
}

func (self *PacketWriter) WritePacketHeader(pktHeader *PacketHeader) error {

	writer := self.binaryStreamWriter

	/* FTS-0001: OrgNode: Word: Origination node address */
	err1 := writer.WriteUINT16(pktHeader.OrigNode)
	if err1 != nil {
		return err1
	}

	/* FTS-0001: DstNode: Word: Destination node address */
	err2 := writer.WriteUINT16(pktHeader.DestNode)
	if err2 != nil {
		return err2
	}

	/* FTS-0001: Year: Word: Year packet generated */
	err3 := writer.WriteUINT16(pktHeader.Year)
	if err3 != nil {
		return err3
	}

	/* FTS-001: Month: Word: --//-- */
	err4 := writer.WriteUINT16(pktHeader.Month)
	if err4 != nil {
		return err4
	}

	/* FTS-001: Day: Word: --//-- */
	err5 := writer.WriteUINT16(pktHeader.Day)
	if err5 != nil {
		return err5
	}

	/* FTS-001: Hour: Word: --//-- */
	err6 := writer.WriteUINT16(pktHeader.Hour)
	if err6 != nil {
		return err6
	}

	/* FTS-001: Min: Word: --//-- */
	err7 := writer.WriteUINT16(pktHeader.Minute)
	if err7 != nil {
		return err7
	}

	/* FTS-001: Sec: Word: --//-- */
	err8 := writer.WriteUINT16(pktHeader.Second)
	if err8 != nil {
		return err8
	}

	/* FTS-001: Baud: Int: Baud Rate (not in use) */
	err9 := writer.WriteUINT16(0)
	if err9 != nil {
		return err9
	}

	/* FTS-0001: PktVer: Int: Packet Version (Always 2) */
	var pktVersion uint16 = 2
	err10 := writer.WriteUINT16(pktVersion)
	if err10 != nil {
		return err10
	}

	/* FTS-0001: OrgNet: Word: Origination net address */
	err11 := writer.WriteUINT16(pktHeader.OrigNet)
	if err11 != nil {
		return err11
	}

	/* FTS-0001: DstNet: Word: Destination net address */
	err12 := writer.WriteUINT16(pktHeader.DestNet)
	if err12 != nil {
		return err12
	}

	/* FTS-0001: PrdCodL: Bytes: FTSC Product Code */
	err13 := writer.WriteUINT8(0)
	if err13 != nil {
		return err13
	}

	/* FTS-0001: PVMajor: Bytes: FTSC Product Rev */
	err14 := writer.WriteUINT8(0)
	if err14 != nil {
		return err14
	}

	/* FTS-0001: Password: 8*Bytes: Packet password (8 byte) */
	pktPassword := make([]byte, 8)
	copy(pktPassword, pktHeader.PktPassword)
	err15 := writer.WriteBytes(pktPassword)
	if err15 != nil {
		return err15
	}

	/* FSC-0039: QOrgZone: Int: Orig Zone */
	err16 := writer.WriteUINT16(pktHeader.OrigZone)
	if err16 != nil {
		return err16
	}

	/* FSC-0039: QDstZone: Int: Dest Zone */
	err17 := writer.WriteUINT16(pktHeader.DestZone)
	if err17 != nil {
		return err17
	}

	/* FSC-0039: Filler: 2*Bytes: Spare Change */
	filler := make([]byte, 2)
	err18 := writer.WriteBytes(filler)
	if err18 != nil {
		return err18
	}

	/* FSC-0039: CapValid: Word: CW Bytes-Swapped Valid Copy */
	var CapValid uint16 = 0x100
	err19 := writer.WriteUINT16(CapValid)
	if err19 != nil {
		return err19
	}

	/* FSC-0039: PrdCodH: Bytes: FTSC Product Code */
	var productCode uint8 = 0
	err20 := writer.WriteUINT8(productCode)
	if err20 != nil {
		return err20
	}

	/* FSC-0039: PVMinor: Bytes: FTSC Product Rev */
	var revision uint8 = 0
	err21 := writer.WriteUINT8(revision)
	if err21 != nil {
		return err21
	}

	/* FSC-0039: CapWord: Word: Capability Word */
	var CapWord uint16 = 1
	err22 := writer.WriteUINT16(CapWord)
	if err22 != nil {
		return err22
	}

	/* FSC-0039: OrigZone: Int: Origination Zone */
	err23 := writer.WriteUINT16(pktHeader.OrigZone)
	if err23 != nil {
		return err23
	}

	/* FSC-0039: DestZone: Int: Destination Zone */
	err24 := writer.WriteUINT16(pktHeader.DestZone)
	if err24 != nil {
		return err24
	}

	/* FSC-0039: OrigPoint: Int: Origination Point */
	err25 := writer.WriteUINT16(pktHeader.OrigPoint)
	if err25 != nil {
		return err25
	}

	/* FSC-0039: DestPoint: Int: Destination Point */
	err26 := writer.WriteUINT16(pktHeader.DestPoint)
	if err26 != nil {
		return err26
	}

	/* FSC-0039: ProdData: Long: Product-specific data */
	var ProdData int32 = 0
	err27 := writer.WriteINT32(ProdData)
	if err27 != nil {
		return err27
	}

	return nil
}

func (self *PacketWriter) WritePackedMessage(packedMessage *PackedMessage) error {

	/* FTS-0001: MsgType: Word: Message Type, old type-1 obsolete */
	var MsgType uint16 = 2
	err1 := self.binaryStreamWriter.WriteUINT16(MsgType)
	if err1 != nil {
		return err1
	}

	/* FTS-0001: OrigNode: Word: Origination node */
	err2 := self.binaryStreamWriter.WriteUINT16(packedMessage.OrigAddr.Node)
	if err2 != nil {
		return err2
	}

	/* FTS-0001: DestNode: Word: Destination node */
	err3 := self.binaryStreamWriter.WriteUINT16(packedMessage.DestAddr.Node)
	if err3 != nil {
		return err3
	}

	/* FTS-0001: OrigNet: Word: Origination net address */
	err4 := self.binaryStreamWriter.WriteUINT16(packedMessage.OrigAddr.Net)
	if err4 != nil {
		return err4
	}

	/* FTS-0001: DestNet: Word: Destination net address */
	err5 := self.binaryStreamWriter.WriteUINT16(packedMessage.DestAddr.Net)
	if err5 != nil {
		return err5
	}

	/* FTS-0001: AttributeWord: Word: Attribute */
	err6 := self.binaryStreamWriter.WriteUINT16(packedMessage.Attributes)
	if err6 != nil {
		return err6
	}

	/* FTS-0001: Cost: Word: ... */
	err7 := self.binaryStreamWriter.WriteUINT16(0)
	if err7 != nil {
		return err7
	}

	/* FTS-0001: DateTime: 20*Bytes: Message body was last edited */
	newFidoDate := fidotime.NewFidoDate()
	newFidoDate.SetNow()
	var DateTime []byte = newFidoDate.FTSC()
	var newDateTime = make([]byte, 20)
	copy(newDateTime, DateTime)
	err8 := self.binaryStreamWriter.WriteBytes(newDateTime)
	if err8 != nil {
		return err8
	}

	/* FTS-0001: ToUserName: N*Bytes: Maximum 36 */
	err9 := self.binaryStreamWriter.WriteZStringWithLimit([]byte(packedMessage.ToUserName), 36-1)
	if err9 != nil {
		return err9
	}

	/* FTS-0001: FromUserName: N*Bytes: Maximum 36 */
	err10 := self.binaryStreamWriter.WriteZStringWithLimit([]byte(packedMessage.FromUserName), 36-1)
	if err10 != nil {
		return err10
	}

	/* FTS-0001: Subject: N*Bytes: Maximum 72 */
	err11 := self.binaryStreamWriter.WriteZStringWithLimit([]byte(packedMessage.Subject), 72-1)
	if err11 != nil {
		return err11
	}

	/* FTS-0001: Text: ... */
	err12 := self.binaryStreamWriter.WriteZString([]byte(packedMessage.Text))
	if err12 != nil {
		return err12
	}

	return nil
}

const (
	SOH = "\x01"
	CR  = "\x0D"
)

func (self *PacketWriter) WritePacketEnd() error {
	var packetEndMarker uint16 = 0
	err1 := self.binaryStreamWriter.WriteUINT16(packetEndMarker)
	if err1 != nil {
		return err1
	}
	return nil
}
