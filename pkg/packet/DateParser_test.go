package packet

import (
	"testing"
	"time"
)

func TimeEqual(t1 time.Time, t2 time.Time) bool {
	return t1.Equal(t2)
}

func TestDateParse(t *testing.T) {

	moscow, err0 := time.LoadLocation("Europe/Moscow")
	if err0 != nil {
		t.Error("Fail load location Europe/Moscow")
		return
	}

	expected := time.Date(2019, time.December, 1, 9, 3, 20, 0, moscow)

	var datetime []byte = []byte("01 Dec 19  09:03:20")

	parser := NewDateParser()

	date, err1 := parser.Parse(datetime)
	if err1 != nil {
		t.Error("Error during parse date")
		return
	}
	if date == nil {
		t.Error("No result provided by parser date")
		return
	}

	actual, err2 := date.Time()
	if err2 != nil {
		t.Error("Error during convert date to time")
		return
	}

	if TimeEqual(*actual, expected) {
	} else {
		t.Errorf("Wrong result in parse date: actual = %q expected = %q", *actual, expected)
		return
	}

}
