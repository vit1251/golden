package packet

import (
	"fmt"
	"testing"
)

func TestMessageBody_Bytes(t *testing.T) {

	msgBody := NewMessageBody()
	msgBody.SetArea("RU.GOLDEN")
	msgBody.SetContent([]byte("Hello, wrold!"))
	var chrsKludge string = "CP866"
	msgBody.AddKludge(Kludge{
		Name: "CHRS",
		Value: chrsKludge,
		Raw: []byte(fmt.Sprintf("\x01CHRS: %s", chrsKludge)),
	})

	newBody := msgBody.Bytes()

	t.Logf("msgBody = %q", newBody)

}
