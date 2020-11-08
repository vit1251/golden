package packet

import "time"

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
	netAddr := NewNetAddr()
	netAddr.SetAddr(addr)

	/* Set Orig address */
	self.DestZone = netAddr.Zone
	self.DestNet = netAddr.Net
	self.DestNode = netAddr.Node
	self.DestPoint = netAddr.Point

}

func (self *PacketHeader) SetDate(newDate time.Time) {
	/* Populate date */
	self.Year = uint16(newDate.Year())
	self.Month = uint16(newDate.Month() - 1)
	self.Day = uint16(newDate.Day())

	/* Populate time */
	self.Hour = uint16(newDate.Hour())
	self.Minute = uint16(newDate.Minute())
	self.Second = uint16(newDate.Second())
}

func (self PacketHeader) GetDate() time.Time {
	var newYear int = int(self.Year)
	var newMonth time.Month = time.Month(self.Month + 1)
	var newDay int = int(self.Day)
	var newHour int = int(self.Hour)
	var newMinute int = int(self.Minute)
	var newSecond int = int(self.Second)
	return time.Date(newYear, newMonth, newDay, newHour, newMinute, newSecond, 0, time.Local)
}
