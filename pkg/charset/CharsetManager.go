package charset

import (
	"golang.org/x/text/encoding/charmap"
)

type CharsetManager struct {

}

func NewCharsetManager() *CharsetManager {
	cm := new(CharsetManager)
	return cm
}

func (self *CharsetManager) DecodeText(source []byte) ([]rune, error) {

	var result []rune

	charmap := charmap.CodePage866

	for _, ch := range source {
		r := charmap.DecodeByte(ch)
		result = append(result, r)
	}

	return result, nil
}

func (self *CharsetManager) EncodeText(source []rune) ([]byte, error) {

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