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
