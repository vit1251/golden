package packet

import "testing"

func TestMsgAttr(t *testing.T) {

	if MsgAttrPrivate == 1 {
		// ignore
	} else {
		t.Errorf("Fail on MsgAttrPrrivate")
	}

	if MsgAttrCrash == 2 {
		// ignore
	} else {
		t.Errorf("Fail on MsgAttrCrash")
	}

}
