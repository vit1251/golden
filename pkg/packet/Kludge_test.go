package packet

import "testing"

func TestKludge_Set(t *testing.T) {
	k := NewKludge()
	row := []byte{'\x01', 'C', 'H', 'R', 'S', ' ', 'C', 'P', '8', '6', '6', ' ', '2'}
	k.Set(row)
	t.Logf("raw = %s", k.Raw)
	t.Logf("name = %q", k.Name)
	t.Logf("value = %q", k.Value)
}

func TestNewKludgeParse(t *testing.T) {
	k := NewKludge()

	k.Set([]byte("\x01INTL 2:5030/1002 2:5030:1001"))
	t.Logf("name = %q value = %q", k.Name, k.Value)

	k.Set([]byte("\x01MSGID: 2:5030/1003.4 01020304"))
	t.Logf("name = %q value = %q", k.Name, k.Value)

}