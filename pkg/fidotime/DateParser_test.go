package fidotime

import (
	"log"
	"testing"
	"time"
)

func TestDateParse(t *testing.T) {

	var datetime []byte = []byte("21 Feb 20  05:18:17")

	parser := NewDateParser()
	date, err1 := parser.Parse(datetime)
	if err1 != nil {
		t.Errorf("Error during parse date: err = %+v", err1)
		return
	}

	if date == nil {
		t.Errorf("No result provided by parser date")
		return
	}

	/* Check date = 21 Feb 20 */
	if date.day != 21 {
		t.Errorf("Wrong day: expect = 21 actual = %d", date.day)
		return
	}
	if date.month != 2 {
		t.Errorf("Wrong month: expect = 2 actual = %d", date.month)
		return
	}
	if date.year != 2020 {
		t.Errorf("Wrong year: expect = 2020 actual = %d", date.year)
		return
	}

	/* Check time = 05:18:17 */
	if date.hour != 5 {
		t.Errorf("Wrong hour: expect = 5 actual = %d", date.hour)
		return
	}
	if date.minute != 18 {
		t.Errorf("Wrong minute: expect = 18 actual = %d", date.minute)
		return
	}
	if date.sec != 17 {
		t.Errorf("Wrong sec: expect = 17 actual = %d", date.sec)
		return
	}

}

func TestDateParseWithZone(t *testing.T) {

	var datetime []byte = []byte("21 Feb 20  05:18:17")

	parser := NewDateParser()
	date, err1 := parser.Parse(datetime)
	if err1 != nil {
		t.Errorf("Wrong parse time: err = %+v", err1)
		return
	}

	zone := time.FixedZone("+0700", +7*60*60)
	newDate, err2 := date.CreateTime(zone)
	if err2 != nil {
		t.Errorf("Wrong create time: err = %+v", err2)
		return
	}
	log.Printf("newDate = %+v", newDate)
	log.Printf("newDateLocal = %+v", newDate.Local())

}
