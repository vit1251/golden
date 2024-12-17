package packet

import (
	"errors"
	"github.com/vit1251/golden/pkg/fidotime"
	"io"
	"log"
)

type PacketReader struct {
	sourceStream       io.Reader
	binaryStreamReader *BinaryReader
}

func NewPacketReader(stream io.Reader) *PacketReader {

	/* Create packet reader */
	pr := new(PacketReader)
	pr.sourceStream = stream

	/* Create binary stream reader */
	binaryStreamReader := NewBinaryReader(stream)
	pr.binaryStreamReader = binaryStreamReader

	return pr
}

func (self *PacketReader) ReadPacketHeader() (*PacketHeader, error) {

	reader := self.binaryStreamReader

	/* Create packet header */
	pktHeader := new(PacketHeader)

	/* FTS-0001: OrgNode: Word: Origination node address */
	OrgNode, err1 := reader.ReadUINT16()
	if err1 != nil {
		return nil, err1
	}
	pktHeader.OrigNode = OrgNode

	/* FTS-0001: DstNode: Word: Destination node address */
	DstNode, err2 := reader.ReadUINT16()
	if err2 != nil {
		return nil, err2
	}
	pktHeader.DestNode = DstNode

	/* FTS-0001: Year: Int: Year packet generated */
	Year, err3 := reader.ReadUINT16()
	if err3 != nil {
		return nil, err3
	}
	pktHeader.Year = Year

	/* FTS-0001: Month: Int: --//-- */
	Month, err4 := reader.ReadUINT16()
	if err4 != nil {
		return nil, err4
	}
	pktHeader.Month = Month + 1

	/* FTS-0001: Day: Int: --//-- */
	Day, err5 := reader.ReadUINT16()
	if err5 != nil {
		return nil, err5
	}
	pktHeader.Day = Day

	/* FTS-0001: Hour: Int: --//-- */
	Hour, err6 := reader.ReadUINT16()
	if err6 != nil {
		return nil, err6
	}
	pktHeader.Hour = Hour

	/* FTS-0001: Min: Int: --//-- */
	Min, err7 := reader.ReadUINT16()
	if err7 != nil {
		return nil, err7
	}
	pktHeader.Minute = Min

	/* FTS-0001: Sec: Int: --//-- */
	Sec, err8 := reader.ReadUINT16()
	if err8 != nil {
		return nil, err8
	}
	pktHeader.Second = Sec

	/* FTS-0001: Baud: Int: Baud Rate (not in use) */
	_, err9 := reader.ReadUINT16()
	if err9 != nil {
		return nil, err9
	}

	/* FTS-0001: PktVer: Int: Packet Version (Always 2) */
	pktVersion, err10 := reader.ReadUINT16()
	if err10 != nil {
		return nil, err10
	}
	if pktVersion != 2 {
		return nil, errors.New("invalid packet version")
	}

	/* FTS-0001: OrgNet: Word: Origination net address */
	OrgNet, err11 := reader.ReadUINT16()
	if err11 != nil {
		return nil, err11
	}
	pktHeader.OrigNet = OrgNet

	/* FTS-0001: DstNet: Word: Destination net address */
	DstNet, err12 := reader.ReadUINT16()
	if err12 != nil {
		return nil, err12
	}
	pktHeader.DestNet = DstNet

	/* FTS-0001: PrdCodL: Bytes: FTSC Product Code */
	_, err13 := reader.ReadUINT8()
	if err13 != nil {
		return nil, err13
	}

	/* FTS-0001: PVMajor: Bytes: FTSC Product Rev */
	_, err14 := reader.ReadUINT8()
	if err14 != nil {
		return nil, err14
	}

	/* FTS-0001: Password: Packet password (8 byte) */
	Password, err15 := reader.ReadBytes(8)
	if err15 != nil {
		return nil, err15
	}
	pktHeader.PktPassword = Password

	/* FSC-0039: QOrgZone: Int: Orig Zone */
	QOrgZone, err16 := reader.ReadINT16()
	if err16 != nil {
		return nil, err11
	}

	/* FSC-0039: QDstZone: Int: Dest Zone */
	QDstZone, err17 := reader.ReadINT16()
	if err17 != nil {
		return nil, err17
	}

	/* FTS-0001: Filler: 2*Bytes: Spare Change */
	_, err18 := reader.ReadBytes(2)
	if err18 != nil {
		return nil, err18
	}

	/* FSC-0039: CapValid: CW Bytes-Swapped Valid Copy */
	CapValid, err19 := reader.ReadUINT16()
	if err19 != nil {
		return nil, err19
	}

	/* FSC-0039: PrdCodH: Bytes: FTSC Product Code */
	_, err20 := reader.ReadUINT8()
	if err20 != nil {
		return nil, err20
	}

	/* FSC-0039: PVMinor: Bytes: FTSC Product Rev */
	_, err21 := reader.ReadUINT8()
	if err21 != nil {
		return nil, err21
	}

	/* FSC-0039: CapWord: Word: Capability Word */
	CapWord, err22 := reader.ReadUINT16()
	if err22 != nil {
		return nil, err22
	}

	/* FSC-0039: OrigZone: Int: Origination Zone */
	OrigZone, err23 := reader.ReadINT16()
	if err23 != nil {
		return nil, err23
	}

	/* FSC-0039: DestZone: Int: Destination Zone */
	DestZone, err24 := reader.ReadINT16()
	if err24 != nil {
		return nil, err24
	}

	/* FSC-0039: OrigPoint: Int: Origination Point */
	OrigPoint, err25 := reader.ReadINT16()
	if err25 != nil {
		return nil, err25
	}

	/* FSC-0039: DestPoint: Int: Destination point */
	DestPoint, err26 := reader.ReadINT16()
	if err26 != nil {
		return nil, err26
	}

	/* FSC-0039: ProdData: Long: Product-specific data */
	_, err27 := reader.ReadINT32()
	if err27 != nil {
		return nil, err27
	}

	/* FSC-0039: Check capability */

	var swapCapValid uint16 = ((CapValid >> 8) + (CapValid << 8)) & 0xFFFF
	if (CapWord != 0) && (CapWord == swapCapValid) {
		log.Printf("Packet process in FSC-0039 capability mode")
	} else {
		log.Printf("Packet process in FSC-0001 capability mode")
		return pktHeader, nil
	}

	log.Printf("QOrgZone = %d QDstZone = %d OrigZone = %d DestZone = %d", QOrgZone, QDstZone, OrigZone, DestZone)

	pktHeader.OrigZone = uint16(QOrgZone)
	if pktHeader.OrigZone == 0 {
		pktHeader.OrigZone = uint16(OrigZone)
	}

	pktHeader.DestZone = uint16(QDstZone)
	if pktHeader.DestZone == 0 {
		pktHeader.DestZone = uint16(DestZone)
	}

	pktHeader.OrigPoint = uint16(OrigPoint)
	pktHeader.DestPoint = uint16(DestPoint)

	return pktHeader, nil
}

func (self *PacketReader) ReadPackedMessage() (*PackedMessage, error) {

	packedMessage := new(PackedMessage)

	/* FTS-0001: MsgType: Word: Message Type, old type-1 obsolete */
	MsgType, err1 := self.binaryStreamReader.ReadUINT16()
	if err1 != nil {
		return nil, err1
	}
	if MsgType == 0 {
		return nil, io.EOF
	}
	if MsgType != 2 {
		return nil, errors.New("invalid packet message version")
	}

	/* FTS-0001: OrigNode: Word: Origination node */
	OrigNode, err2 := self.binaryStreamReader.ReadUINT16()
	if err2 != nil {
		return nil, err2
	}
	packedMessage.OrigAddr.Node = OrigNode

	/* FTS-0001: DestNode: Word: Destination node */
	DestNode, err3 := self.binaryStreamReader.ReadUINT16()
	if err3 != nil {
		return nil, err3
	}
	packedMessage.DestAddr.Node = DestNode

	/* FTS-0001: OrigNet: Word: Origination net address */
	OrigNet, err4 := self.binaryStreamReader.ReadUINT16()
	if err4 != nil {
		return nil, err4
	}
	packedMessage.OrigAddr.Net = OrigNet

	/* FTS-0001: DestNet: Word: Destination net address */
	DestNet, err5 := self.binaryStreamReader.ReadUINT16()
	if err5 != nil {
		return nil, err5
	}
	packedMessage.DestAddr.Net = DestNet

	/* FTS-0001: AttributeWord: Word: Attribute */
	AttributeWord, err6 := self.binaryStreamReader.ReadUINT16()
	if err6 != nil {
		return nil, err6
	}
	packedMessage.Attributes = AttributeWord

	/* FTS-0001: Cost: Word: ... */
	_, err7 := self.binaryStreamReader.ReadUINT16()
	if err7 != nil {
		return nil, err7
	}

	/* FTS-0001: DateTime: 20*Bytes: Message body was last edited */
	DateTime, err8 := self.binaryStreamReader.ReadBytes(20)
	if err8 != nil {
		return nil, err8
	}
	log.Printf("PakdMessage: DateTime = %s", DateTime)

	/* Create new one parser */
	parser := fidotime.NewDateParser()
	if newDateTime, err := parser.Parse(DateTime); err != nil {
		return nil, err
	} else {
		packedMessage.Time = newDateTime
	}

	/* FTS-0001: ToUserName: 36*Bytes:  */
	ToUserName, err9 := self.binaryStreamReader.ReadZString()
	if err9 != nil {
		return nil, err9
	}
	packedMessage.ToUserName = ToUserName

	/* FTS-0001: FromUserName: 36*Bytes: */
	FromUserName, err10 := self.binaryStreamReader.ReadZString()
	if err10 != nil {
		return nil, err10
	}
	packedMessage.FromUserName = FromUserName

	/* FTS-0001: Subject: 72*Bytes */
	Subject, err11 := self.binaryStreamReader.ReadZString()
	if err11 != nil {
		return nil, err11
	}
	packedMessage.Subject = Subject

	/* FTS-0001: Text: ... */
	Text, err1 := self.binaryStreamReader.ReadZString()
	if err1 != nil {
		return nil, err1
	}
	packedMessage.Text = Text

	return packedMessage, nil
}
