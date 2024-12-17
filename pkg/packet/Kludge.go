package packet

import "bytes"

type Kludge struct {
	Name  string
	Value string
	Raw   []byte
}

func NewKludge() *Kludge {
	return new(Kludge)
}

type KludgeState int8

const (
	KludgeStateName  KludgeState = 1
	KludgeStateValue KludgeState = 2
)

func (self *Kludge) parseKludge(kludge []byte) {
	name := ""
	value := ""
	state := KludgeStateName
	for _, ch := range kludge {
		if state == KludgeStateName {
			if ch == ' ' || ch == ':' {
				state = KludgeStateValue
			} else {
				name += string(ch)
			}
		} else if state == KludgeStateValue {
			value += string(ch)
		}
	}

	/* Setup kludge */
	self.Name = name
	self.Value = value

}

func (self *Kludge) Set(raw []byte) {

	/* Save RAW value */
	self.Raw = raw

	/* Process */
	if bytes.HasPrefix(raw, []byte{'\x01'}) {
		content := bytes.TrimPrefix(raw, []byte{'\x01'})
		self.parseKludge(content)
	}

}
