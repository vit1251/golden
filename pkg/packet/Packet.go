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
	OrigAddr         NetAddr
	DestAddr         NetAddr
	pktCreated       PktDateTime
	capatiblityByte1 byte
	capatiblityByte2 byte
	hiProductCode    byte
	minorProductRev  byte
	capabilityWord   uint16
	loProductCode    byte
	majorProductRev  byte
}

func NewPacketHeader() (*PacketHeader) {
	ph := new(PacketHeader)
	return ph
}

type PacketMessage struct {
	OrigAddr      NetAddr
	DestAddr      NetAddr
	Attributes    uint16
	ToUserName    string
	FromUserName  string
	Subject       string
	Text          string
	RAW         []byte
	Time         *time.Time   /* Packet time */
}

func NewPacketMessage() (*PacketMessage) {
	pm := new(PacketMessage)
	return pm
}

func (self *PacketMessage) UnsetAttribute(attribute string) (error) {
	return nil
}

func (self *PacketMessage) SetAttribute(attribute string) (error) {
	return nil
}

func (self *PacketMessage) SetToUserName(ToUserName string) (error) {
	return nil
}

func (self *PacketMessage) SetFromUserName(FromUserName string) (error) {
	return nil
}

func (self *PacketMessage) SetSubject(subj string) (error) {
	return nil
}

func (self *PacketMessage) SetText(msg string) (error) {
	return nil
}

func (self *PacketMessage) SetTime(t *time.Time) (error) {
	return nil
}
