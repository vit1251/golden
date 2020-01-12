package ftn

import (
	"bytes"
	"log"
)

type NetAddressParserState int

const (

	UnknownState  NetAddressParserState = 0

	ZoneState     NetAddressParserState = 1
	NetState      NetAddressParserState = 2
	NodeState     NetAddressParserState = 3
	PointState    NetAddressParserState = 4

)

type NetAddressParser struct {
	state     NetAddressParserState   /* Parser state    */
	addr      NetAddress              /* Result address  */
	cache     bytes.Buffer            /* Cache           */
}

func NewNetAddressParser() (*NetAddressParser) {
	ap := new(NetAddressParser)
	ap.state = ZoneState
	return ap
}

func (self *NetAddressParser) processComplete() {

	if self.state == ZoneState {
		self.addr.Zone = self.cache.String()
		self.cache.Reset()
		self.state = NetState
	} else if self.state == NetState {
		self.addr.Net = self.cache.String()
		self.cache.Reset()
		self.state = NodeState
	} else if self.state == NodeState {
		self.addr.Node = self.cache.String()
		self.cache.Reset()
		self.state = PointState
	} else if self.state == PointState {
		self.addr.Point = self.cache.String()
		self.cache.Reset()
		self.state = UnknownState
	}

}

func (self *NetAddressParser) Parse(addr string) (*NetAddress, error) {

	for _, rune := range addr {
		if self.state == ZoneState {
			if rune == ':' {
				self.processComplete()
			} else {
				self.cache.WriteRune(rune)
			}
		} else if self.state == NetState {
			if rune == '/' {
				self.processComplete()
			} else {
				self.cache.WriteRune(rune)
			}
		} else if self.state == NodeState {
			if rune == '.' {
				self.processComplete()
			} else {
				self.cache.WriteRune(rune)
			}
		} else if self.state == PointState {
			self.cache.WriteRune(rune)
		}
	}

	/* Last section */
	self.processComplete()

	/* Debug message */
	log.Printf("NetAddressParser.Parse: %s => %+v", addr, self.addr)

	/* Done */
	return &self.addr, nil
}
