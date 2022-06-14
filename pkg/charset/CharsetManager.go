package charset

import (
	"fmt"
	"github.com/vit1251/golden/pkg/registry"
	"golang.org/x/text/encoding/charmap"
	"unicode/utf8"
)

type CharsetManager struct {
	mapping map[string]*charmap.Charmap
}

func NewCharsetManager(r *registry.Container) *CharsetManager {
	cm := new(CharsetManager)
	cm.mapping = make(map[string]*charmap.Charmap)
	//
	cm.registerCharset("CP866", charmap.CodePage866)
	cm.registerCharset("LATIN-1", charmap.ISO8859_1)
	cm.registerCharset("CP437", charmap.CodePage437)
	cm.registerCharset("ASCII", charmap.ISO8859_1) // TODO - processing as special case ...
	//
	return cm
}

func (self *CharsetManager) registerCharset(name string, c *charmap.Charmap) {
	self.mapping[name] = c
}

func (self *CharsetManager) decode(source []byte, charmap *charmap.Charmap) ([]rune, error) {

	var result []rune

	for _, ch := range source {
		r := charmap.DecodeByte(ch)
		result = append(result, r)
	}

	return result, nil
}

func (self *CharsetManager) encode(source []rune, charmap *charmap.Charmap) ([]byte, error) {

	var result []byte

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

	var convertationMap *charmap.Charmap = self.searchCharmapByIndex(charset)

	if convertationMap != nil {
		var result string
		if unicodeBody, err1 := self.decode(msgBody, convertationMap); err1 == nil {
			result = string(unicodeBody)
		} else {
			return result, err1
		}
		return result, nil
	} else if charset == "UTF-8" {
		var result string
		result = string(msgBody)
		return result, nil
	}

	return "", fmt.Errorf("wrong charset on message")
}

func (self CharsetManager) EncodeMessageBody(msgBody string, charset string) ([]byte, error) {

	var convertationMap *charmap.Charmap = self.searchCharmapByIndex(charset)

	if convertationMap != nil {
		var newMsgBody []rune = []rune(msgBody)
		if unicodeBody, err1 := self.encode(newMsgBody, convertationMap); err1 == nil {
			return unicodeBody, nil
		} else {
			return nil, err1
		}
	} else if charset == "UTF-8" {

		var result []byte
		for _, r := range msgBody {
			buf := make([]byte, 4)
			n := utf8.EncodeRune(buf, r)
			buf = buf[:n]
			result = append(result, buf...)
		}
		return result, nil

	}

	return nil, nil

}

func (self CharsetManager) searchCharmapByIndex(charmapIndex string) *charmap.Charmap {
	if value, ok := self.mapping[charmapIndex]; ok {
		return value
	}
	return nil
}
