package packet

type PacketHeader struct {

	/* Origination */
	OrigZone     uint16
	OrigNet      uint16
	OrigNode     uint16
	OrigPoint    uint16

	/* Destination */
	DestZone     uint16
	DestNet      uint16
	DestNode     uint16
	DestPoint    uint16

	/* Date and time */
	Year         uint16
	Month        uint16
	Day          uint16

	Hour         uint16
	Minute       uint16
	Second       uint16

	/* Password */
	PktPassword      []byte

}

func (self *PacketHeader) SetOrigAddr(addr string) {

	/* Parse address */
	netAddr := NetAddr{}
	netAddr.SetAddr(addr)

	/* Set Orig address */
	self.OrigZone = netAddr.Zone
	self.OrigNet = netAddr.Net
	self.OrigNode = netAddr.Node
	self.OrigPoint = netAddr.Point

}

func (self *PacketHeader) SetDestAddr(addr string) {

	/* Parse address */
	netAddr := NetAddr{}
	netAddr.SetAddr(addr)

	/* Set Orig address */
	self.DestZone = netAddr.Zone
	self.DestNet = netAddr.Net
	self.DestNode = netAddr.Node
	self.DestPoint = netAddr.Point

}
