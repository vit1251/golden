package packet

import (
	"fmt"
	"time"
)

type NetAddr struct {
	Zone    uint16
	Net     uint16
	Node    uint16
	Point   uint16
}

func (self *NetAddr) SetAddr(addr string) (error) {
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

type PacketHeader struct {
	OrigAddr            NetAddr
	DestAddr            NetAddr
	pktCreated          PktDateTime
	capatiblityByte1    uint8
	capatiblityByte2    uint8
	pktPassword       []byte
	auxNet              uint16
	hiProductCode       uint8
	minorProductRev     uint8
	capabilityWord      uint16
	loProductCode       uint8
	majorProductRev     uint8
}

func NewPacketHeader() (*PacketHeader) {
	ph := new(PacketHeader)
	return ph
}

func (self PacketHeader) IsCapatiblity() (bool) {
	var capWord uint16 = uint16(self.capatiblityByte1 << 8) + uint16(self.capatiblityByte2)
	return capWord == self.capabilityWord
}

type PacketMessageHeader struct {
	OrigAddr      NetAddr
	DestAddr      NetAddr
	Attributes    uint16
	ToUserName    string
	FromUserName  string
	Subject       string
	Time         *time.Time   /* Packet time */
}

func NewPacketMessageHeader() (*PacketMessageHeader) {
	msgHeader := new(PacketMessageHeader)
	return msgHeader
}

func (self *PacketMessageHeader) UnsetAttribute(attribute string) (error) {
	return nil
}

func (self *PacketMessageHeader) SetAttribute(attribute string) (error) {
	return nil
}

func (self *PacketMessageHeader) SetToUserName(ToUserName string) (error) {
	return nil
}

func (self *PacketMessageHeader) SetFromUserName(FromUserName string) (error) {
	return nil
}

func (self *PacketMessageHeader) SetSubject(subj string) (error) {
	return nil
}

func (self *PacketMessageHeader) SetTime(t *time.Time) (error) {
	return nil
}
