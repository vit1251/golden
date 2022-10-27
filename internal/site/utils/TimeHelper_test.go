package utils

import (
	"testing"
)

// Green path
func Test1_TimeHelper_makeDuration(t *testing.T) {

	var source uint64 = (((((2 * 60) + 30) * 60) + 25) * 1000) + 123

	actual := TimeHelper_renderDurationInMilli(source)
	expected := "02:30:25.123"

	if actual != expected {
		t.Fatalf("Test1 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}
