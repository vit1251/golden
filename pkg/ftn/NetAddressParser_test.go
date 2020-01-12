package ftn

import (
	"testing"
)

func TestNetAddressParser1(t *testing.T) {

	nap := NewNetAddressParser()
	addr, err := nap.Parse("2:5030/1592.11")
	if err != nil {
		t.Errorf("Fail on Parse with error: err = %q", err)
	}

	if addr.Zone != "2" {
		t.Errorf("Wrong zone: addr = %q", addr)
	}

	if addr.Net != "5030" {
		t.Errorf("Wrong net: addr = %q", addr)
	}

	if addr.Node != "1592" {
		t.Errorf("Wrong node: addr = %q", addr)
	}

	if addr.Point != "11" {
		t.Errorf("Wrong point: addr = %q", addr)
	}

}

func TestNetAddressParser2(t *testing.T) {

	nap := NewNetAddressParser()
	addr, err := nap.Parse("2:5030/1592")
	if err != nil {
		t.Errorf("Fail on Parse with error: err = %q", err)
	}

	if addr.Zone != "2" {
		t.Errorf("Wrong zone: addr = %q", addr)
	}

	if addr.Net != "5030" {
		t.Errorf("Wrong net: addr = %q", addr)
	}

	if addr.Node != "1592" {
		t.Errorf("Wrong node: addr = %q", addr)
	}

	if addr.Point != "" {
		t.Errorf("Wrong point: addr = %q", addr)
	}

}
