package squish

import (
	"fmt"
)

type NetAddress struct {
	Zone uint16  // FidoNet zone number
	Net uint16   // FidoNet net number
	Node uint16  // FidoNet node number.
	Point uint16 // FidoNet point number
}

func (self NetAddress) GetAddr() string {
	var result string
	if self.Point == 0 {
		result = fmt.Sprintf("%d:%d/%d", self.Zone, self.Net, self.Node)
	} else {
		result = fmt.Sprintf("%d:%d/%d.%d", self.Zone, self.Net, self.Node, self.Point)
	}
	return result
}
