package fidotime

import (
	"log"
	"testing"
)

func TestTimeZoneParser0(t *testing.T) {
	zp := NewTimeZoneParser()
	zone, err1 := zp.Parse("0300")
	log.Printf("zone = %q", zone)
	if err1 != nil {
		t.Errorf("Wrong timze zone: err1 = %+v", err1)
	}
}

func TestTimeZoneParser1(t *testing.T) {
	zp := NewTimeZoneParser()
	zone, err1 := zp.Parse("-0700")
	log.Printf("zone = %q", zone)
	if err1 != nil {
		t.Errorf("Wrong timze zone: err1 = %+v", err1)
	}
}

func TestTimeZoneParser2(t *testing.T) {
	zp := NewTimeZoneParser()
	zone, err1 := zp.Parse("-0700 ")
	log.Printf("zone = %q", zone)
	if err1 == nil {
		t.Errorf("Wrong timze zone: err1 = %+v", err1)
	}
}

func TestTimeZoneParser3(t *testing.T) {
    zp := NewTimeZoneParser()
	zone, err1 := zp.Parse(" ")
	log.Printf("zone = %q", zone)
	if err1 == nil {
		t.Errorf("Wrong timze zone: err1 = %+v", err1)
	}
}

func TestTimeZoneParser4(t *testing.T) {
	zp := NewTimeZoneParser()
	zone, err1 := zp.Parse("MSK")
	log.Printf("zone = %q", zone)
	if err1 == nil {
		t.Errorf("Wrong timze zone: err1 = %+v", err1)
	}
}
