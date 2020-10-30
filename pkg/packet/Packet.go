package packet

import (
	"fmt"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/ftn"
	"log"
	"strconv"
)

type NetAddr struct {
	Zone    uint16
	Net     uint16
	Node    uint16
	Point   uint16
}

func (self *NetAddr) SetAddr(addr string) error {
	/* Parse address */
	nap := ftn.NewNetAddressParser()
	newAddr, err1 := nap.Parse(addr)
	if err1 != nil {
		return err1
	}
	/* Set address */
	newZone, _ := strconv.Atoi(newAddr.Zone)
	self.Zone = uint16(newZone)
	//
	newNet, _ := strconv.Atoi(newAddr.Net)
	self.Net = uint16(newNet)
	//
	newNode, _ := strconv.Atoi(newAddr.Node)
	self.Node = uint16(newNode)
	//
	newPoint, _ := strconv.Atoi(newAddr.Point)
	self.Point = uint16(newPoint)
	//
	return nil
}

func (self NetAddr) String() string {
	var result string
	if self.Point == 0 {
		result = fmt.Sprintf("%d:%d/%d", self.Zone, self.Net, self.Node)
	} else {
		result = fmt.Sprintf("%d:%d/%d.%d", self.Zone, self.Net, self.Node, self.Point)
	}
	return result
}

type PktDateTime struct {
	Year    uint16
	Mon     uint16
	MDay    uint16
	Hour    uint16
	Min     uint16
	Sec     uint16
}

type PktVersion int

const (
	PKTv2      = 0x0200
	PKTv2plus  = 0x0201
	PKTv22     = 0x0202
)

type PacketHeader struct {
	OrigAddr         NetAddr
	DestAddr         NetAddr
	pktCreated       PktDateTime
	subVersion       uint16
	capatiblityByte1 uint8
	capatiblityByte2 uint8
	pktPassword      []byte
	auxNet           uint16
	hiProductCode    uint8
	minorProductRev  uint8
	capabilityWord   uint16
	loProductCode    uint8
	majorProductRev  uint8
}

func NewPacketHeader() *PacketHeader {
	ph := new(PacketHeader)
	ph.capatiblityByte1 = 0
	ph.capatiblityByte2 = 1
	ph.pktPassword = make([]byte, 8)
	ph.auxNet = 0
	ph.capabilityWord = 1
	return ph
}

func (self *PacketHeader) GetVersion() PktVersion {
	if self.subVersion == 2 {
		return PKTv22
	} else {
		checkFeature := true // TODO - if lo(cap word) == hi(inv cap word) and hi(cap word) == lo( inv cap word) {
		if checkFeature {
			return PKTv2plus
		} else {
			return PKTv2
		}
	}
}

func (self *PacketHeader) SetPassword(password string) {
	newPasswordSize := 8
	newPassword := make([]byte, newPasswordSize)
	copy(newPassword, password)
	log.Printf("newPass: str = %s raw = %+v", newPassword, newPassword)
	self.pktPassword = newPassword
}

func (self PacketHeader) IsCapatiblity() (bool) {
	var capWord uint16 = uint16(self.capatiblityByte1 << 8) + uint16(self.capatiblityByte2)
	return capWord == self.capabilityWord
}

type PacketMessageHeader struct {
	OrigAddr        NetAddr
	DestAddr        NetAddr
	Attributes      uint16
	ToUserName    []byte
	FromUserName  []byte
	Subject       []byte
	Time           *fidotime.FidoDate
}

func NewPacketMessageHeader() *PacketMessageHeader {
	msgHeader := new(PacketMessageHeader)
	return msgHeader
}

type PacketAttr int8

const (
	PacketAttrDirect PacketAttr = 0x01
)

func (self *PacketMessageHeader) UnsetAttribute(attr PacketAttr) error {
	return nil
}

func (self *PacketMessageHeader) SetAttribute(attr PacketAttr) error {
	return nil
}

func (self *PacketMessageHeader) SetToUserName(ToUserName []byte) error {
	self.ToUserName = ToUserName
	return nil
}

func (self *PacketMessageHeader) SetFromUserName(FromUserName []byte) error {
	self.FromUserName = FromUserName
	return nil
}

func (self *PacketMessageHeader) SetSubject(s []byte) error {
	self.Subject = s
	return nil
}

func (self *PacketMessageHeader) SetTime(t *fidotime.FidoDate) error {
	self.Time = t
	return nil
}
