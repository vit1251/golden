package fidotime

import (
	"log"
	"testing"
)

func TestDurationParser(t *testing.T) {

	dp := NewDurationParser()
	d, err1 := dp.Parse("1h10m20s")
	if err1 != nil {
		t.Errorf("Wrong parse string: err1 = %+v", err1)
	}
	log.Printf("duration = %+v", d)

}
