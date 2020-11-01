package packet

import "bytes"

type Kludge struct {
	Name    string
	Value   string
	Raw     []byte
}

func NewKludge() *Kludge {
	return new(Kludge)
}

func (k *Kludge) Set(raw []byte) {

	/* Save RAW value */
	k.Raw = raw

	/* Process */
	if bytes.HasPrefix(raw, []byte{'\x01'}) {
		content := raw[1:]
		parts := bytes.SplitN(content, []byte(" "), 2)
		if len(parts) == 2 {
			k.Name = string(parts[0])
			k.Value = string(parts[1])
		}
	}

}
