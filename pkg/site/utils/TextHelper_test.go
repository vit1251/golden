package utils

import (
    "testing"
)

// Green path
func Test1_TextHelper_makeNameTitle(t *testing.T) {

	source := "Andrey Mundirov"
	
	actual := TextHelper_makeNameTitle(source)
	expected := "AM"

	if actual != expected {
		t.Fatalf("Test1 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}

// Wrong path: tail space
func Test2_TextHelper_makeNameTitle(t *testing.T) {

	source := "Andrey Mundirov "
	
	actual := TextHelper_makeNameTitle(source)
	expected := "AM"

	if actual != expected {
		t.Fatalf("Test2 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}

// Wrong path: head space
func Test3_TextHelper_makeNameTitle(t *testing.T) {

	source := " Andrey Mundirov"
	
	actual := TextHelper_makeNameTitle(source)
	expected := "AM"

	if actual != expected {
		t.Fatalf("Test3 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", actual, expected)
	}

}
