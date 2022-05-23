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
		t.Fatalf("Invalid")
	}

}

// Wrong path: tail space
func Test2_TextHelper_makeNameTitle(t *testing.T) {

	source := "Andrey Mundirov "
	
	actual := TextHelper_makeNameTitle(source)
	expected := "AM"

	if actual != expected {
		t.Fatalf("Invalid")
	}

}

// Wrong path: head space
func Test3_TextHelper_makeNameTitle(t *testing.T) {

	source := " Andrey Mundirov"
	
	actual := TextHelper_makeNameTitle(source)
	expected := "AM"

	if actual != expected {
		t.Fatalf("Invalid")
	}

}
