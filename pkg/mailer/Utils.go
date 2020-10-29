package mailer

import (
	"fmt"
)

type EscapeMode int

const (
	LEGACY EscapeMode = 0
	MODERN EscapeMode = 1
)

func IsAlpha(r rune) bool {
	return false
}

func escape(src string, mode EscapeMode) string {

	var result string
	for _, r := range src {
		if IsAlpha(r) {
			result += fmt.Sprintf("\\x%x", r)
		} else {
			result += fmt.Sprintf("%c", r)
		}
	}
	return result
}

func unescape(src string) (string, error) {
	return src, nil
}
