package tosser

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	m := &TosserManager{}
	newZone := m.makeTimeZone()
	fmt.Printf("value = %v", newZone)
	if newZone != "0300" {
		t.Errorf("Wrong timezone")
	}

}
