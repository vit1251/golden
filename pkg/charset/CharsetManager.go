package charset

import (
	"fmt"
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

func (self CharsetManager) DecodeMessageBody(msgBody []byte, charset string) (string, error) {
	var result string
	if charset == "CP866" {
		if unicodeBody, err1 := self.Decode(msgBody); err1 == nil {
			result = string(unicodeBody)
		} else {
				return result, err1
			}
	} else if charset == "UTF-8" {
		result = string(msgBody)
	} else if charset == "LATIN-1" {
		result = string(msgBody)
	} else {
		return result, fmt.Errorf("wrong charset on message")
	}
	return result, nil
}
