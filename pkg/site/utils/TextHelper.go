package utils

import (
	"fmt"
	"hash/crc32"
	"strings"
)

func TextHelper_makeNameTitle(name string) string {

	parts := strings.Split(name, " ")
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
	crc32value := fmt.Sprintf("%08x", crc32.Checksum([]byte(source), crc32q))
	return fmt.Sprintf("#%s", crc32value[2:8])
}
