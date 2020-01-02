package ftn

import (
	"testing"
//	"time"
)

func TestNetAddressParser(t *testing.T) {

	nap := NewNetAddressParser()
	addr, err := nap.Parse("2:5030/1592.11")
	if err != nil {
		t.Errorf("Fail on Parse with error: err = %q", err)
	}
	t.Errorf("Addr = %q", addr)

}
