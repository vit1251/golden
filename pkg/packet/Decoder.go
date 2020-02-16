package packet

import (
//	"log"
	"golang.org/x/text/encoding/charmap"
)

func DecodeText(source []byte) ([]rune, error) {

	var result []rune

	charmap := charmap.CodePage866

	for _, ch := range source {
		r := charmap.DecodeByte(ch)
		result = append(result, r)
	}

	return result, nil
}

func EncodeText(source []rune) ([]byte, error) {

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