package tosser

import (
	"bytes"
	"log"
	"regexp"
)

type OriginParser struct {
}

func NewOriginParser() *OriginParser {
	return new(OriginParser)
}

func (p OriginParser) Parse(origin []byte) []byte {
	re := regexp.MustCompile(`\([0-9\:\/\.]+\)`)
	addr := re.Find(origin)
	if addr != nil {
		addr = bytes.TrimPrefix(addr, []byte("("))
		addr = bytes.TrimSuffix(addr, []byte(")"))
	}
	log.Printf("addr = %q", addr)
	return addr
}
