package fidotime

import (
	"errors"
	"time"
	"unicode"
)

type DurationParser struct {
}

func NewDurationParser() *DurationParser {
	dp := new(DurationParser)
	return dp
}

func (self *DurationParser) Parse(duration string) (time.Duration, error) {
	var result time.Duration
	var number int = 0
	for _, ch := range duration {
		if unicode.IsDigit(ch) {
			number = 10 * number
			number += int(ch - '0')
		} else {
			if ch == 'm' {
				var newDuration time.Duration = time.Minute * time.Duration(number)
				result = result + newDuration
			} else if ch == 'h' {
				var newDuration time.Duration = time.Hour * time.Duration(number)
				result = result + newDuration
			} else if ch == 's' {
				var newDuration time.Duration = time.Second * time.Duration(number)
				result = result + newDuration
			} else {
				return result, errors.New("problem while parse")
			}
			number = 0
		}
	}

	return result, nil
}
