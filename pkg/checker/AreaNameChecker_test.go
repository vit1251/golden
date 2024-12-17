package checker

import (
	"testing"
)

// Green path
func Test1_1_AreaName_Check(t *testing.T) {

	source := "RU.ANECDOT"

	actual := AreaName_Check(source)
	expected := true

	if actual != expected {
		t.Fatalf("Test1[1] is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}

// Green path: special
func Test1_2_AreaName_Check(t *testing.T) {

	source := "RU.HAM-RADIO"

	actual := AreaName_Check(source)
	expected := true

	if actual != expected {
		t.Fatalf("Test1[2] is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}

// Wrong: russian chars
func Test2_AreaName_Check(t *testing.T) {

	source := "РФ.ОБЬЯСНЯЕМ"

	actual := AreaName_Check(source)
	expected := false

	if actual != expected {
		t.Fatalf("Test2 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}

// Wrong: special chars
func Test3_AreaName_Check(t *testing.T) {

	source := "RU.ANEC\tD\nO\aT"

	actual := AreaName_Check(source)
	expected := false

	if actual != expected {
		t.Fatalf("Test3 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}

// Wrong: puctuation chars
func Test4_AreaName_Check(t *testing.T) {

	source := "RU.AN!E?C&D?OT"

	actual := AreaName_Check(source)
	expected := false

	if actual != expected {
		t.Fatalf("Test4 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}
