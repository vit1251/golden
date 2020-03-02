package fidotime

import (
	"log"
	"testing"
)

// Thanks:
// - Nil Alexandrov for message 5e4f062a for research from 21 Feb 20  05:18:17 in 0700 zone

func TestTimeZoneParser(t *testing.T) {

	zp := NewTimeZoneParser()
	zone, err1 := zp.Parse("0700")
	if err1 != nil {
		t.Errorf("Wrong timze zone: err1 = %+v", err1)
	}
	log.Printf("zone = %+v", zone)

}