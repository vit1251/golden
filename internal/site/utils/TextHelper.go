package utils

import (
	"fmt"
	"hash/crc32"
	"image/color"
	"strings"
)

func TextHelper_makeNameTitle(source string) string {

	/* Stage 1. Smart name splitting */
	var parts []string
	var mem strings.Builder

	for _, ch := range source {
		if ch == ' ' {
			if mem.Len() > 0 {
				part := mem.String()
				parts = append(parts, part)
			}
			mem.Reset()
		} else {
			mem.WriteRune(ch)
		}
	}
	if mem.Len() > 0 {
		part := mem.String()
		parts = append(parts, part)
	}

	/* Stage 2. Construct short name */

	partCount := len(parts)
	if partCount >= 2 {
		firstName := parts[0]
		lastName := parts[partCount-1]
		return fmt.Sprintf("%c%c", firstName[0], lastName[0])
	} else if partCount == 1 {
		firstName := parts[0]
		return fmt.Sprintf("%s", firstName[0:1])
	} else {
		return "?"
	}

}

func TextHelper_makeColorByName(source string) string {

	crc32q := crc32.MakeTable(0xD5828281)
	v := crc32.Checksum([]byte(source), crc32q)

	v1 := byte(0xff & v)
	v2 := byte(0xff & (v >> 8))
	v3 := byte(0xff & (v >> 16))
	v4 := byte(0xff & (v >> 24))

	if v4 < 128 {
		v4 = v4 + 128
	}

	r, g, b := color.CMYKToRGB(v1, v2, v3, v4)
	c := fmt.Sprintf("#%02X%02X%02X", r, g, b)

	return c

}
