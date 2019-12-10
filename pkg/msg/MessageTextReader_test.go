package msg

import (
	"testing"
)

func TestMessage(t *testing.T) {

	var msg string = "Hello, All!\n" +
	    "\n"+
	    " VS> Level1\n" +
	    "\n"+
	    " VS>> Level2\n" +
	    " VS>> Level2\n" +
	    "\n"+
	    " VS>>> Level3\n" +
	    " VS>>> Level3\n" +
	    "\n"+
	    " VS>> Level2\n" +
	    " VS>> Level2\n" +
	    "\n"+
	    " VS> Level1\n" +
	    " VS> Level1\n" +
	    "\n"+
	    "Всего хорошего.\n"

	mr := NewMessageTextReader()
	outDoc := mr.Prepare(msg)

	t.Errorf("outDoc = %s", outDoc)
}
