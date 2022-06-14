package utils

import (
	"testing"
)

// Green path
func Test1_FidoHelper_CreatNetAddr(t *testing.T) {

	got, err1 := CreateNetAddr("2:5030/1592")
	if err1 != nil {
		panic(err1)
	}

	var want string = "f1592.n5030.z2.binkp.net"
	if got != want {
		t.Fatalf("Test1 is FAIL.\n\tACTUAL = %v\n\tEXPECTED = %v", got, want)
	}

}
