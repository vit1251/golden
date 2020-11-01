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
