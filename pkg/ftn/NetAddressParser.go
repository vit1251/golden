package ftn

import (
	"bytes"
	"log"
)

type NetAddressParserState int

const (
	ZoneState     NetAddressParserState = 1
	NetState      NetAddressParserState = 2
	NodeState     NetAddressParserState = 3
	PointState    NetAddressParserState = 4
)

type NetAddressParser struct {
	State     NetAddressParserState   /* Parser state */
	addr      NetAddress
}

func NewNetAddressParser() (*NetAddressParser) {
	ap := new(NetAddressParser)
	ap.State = ZoneState
	return ap
}

func (self *NetAddressParser) Parse(addr string) (*NetAddress, error) {
	var buffer bytes.Buffer
	//
	for _, rune := range addr {
		if self.State == ZoneState {
			if rune == ':' {
				self.addr.Zone = buffer.String()
				buffer.Reset()
				self.State = NetState
			} else {
				buffer.WriteRune(rune)
			}
		} else if self.State == NetState {
			if rune == '/' {
				self.addr.Net = buffer.String()
				buffer.Reset()
				self.State = NodeState
			} else {
				buffer.WriteRune(rune)
			}
		} else if self.State == NodeState {
			if rune == '.' {
				self.addr.Node = buffer.String()
				buffer.Reset()
				self.State = PointState
			} else {
				buffer.WriteRune(rune)
			}
		} else if self.State == PointState {
			buffer.WriteRune(rune)
		}
	}
	//
	if self.State == PointState {
		self.addr.Point = buffer.String()
		buffer.Reset()
	}
	//
	log.Printf("add = %q", self.addr)
	//
	return &self.addr, nil
}
