package charset

import (
	"github.com/vit1251/golden/pkg/registry"
	"golang.org/x/text/encoding/charmap"
)

type CharsetManager struct {
}

func NewCharsetManager(r *registry.Container) *CharsetManager {
	cm := new(CharsetManager)
	return cm
}

func (self *CharsetManager) DecodeString(source []byte) (string, error) {
	var result string
	runes, err := self.Decode(source)
	result = string(runes)
	return result, err
}

func (self *CharsetManager) Decode(source []byte) ([]rune, error) {

	var result []rune

	charmap := charmap.CodePage866

	for _, ch := range source {
		r := charmap.DecodeByte(ch)
		result = append(result, r)
	}

	return result, nil
}

//func (self *CharsetManager) EncodeString(source string) ([]byte, error) {
// TODO - implement it later or newer ...
//}

func (self *CharsetManager) Encode(source []rune) ([]byte, error) {

	var result []byte

	charmap := charmap.CodePage866

	for _, ch := range source {
		if r, ok := charmap.EncodeRune(ch); ok {
			result = append(result, r)
		} else {
			result = append(result, byte('?'))
		}
	}

	return result, nil
}