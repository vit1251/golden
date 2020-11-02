package packet

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ftn"
	"strconv"
)

type NetAddr struct {
	Zone    uint16
	Net     uint16
	Node    uint16
	Point   uint16
}

func NewNetAddr() *NetAddr {
	return new(NetAddr)
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

