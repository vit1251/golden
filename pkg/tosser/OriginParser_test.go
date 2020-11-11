package tosser

import "testing"

func TestOriginParser_Parse(t *testing.T) {

	originParser := NewOriginParser()
	origAddr := originParser.Parse([]byte("Proff (2:5030/1492.11)"))

	t.Logf("origAddr = %s", origAddr)

}
