package utils

import (
	"fmt"
	"time"
)

func DateHelper_monthAsNumber(month time.Month) int {
	if month == time.January {
		return 1
	} else if month == time.February {
		return 2
	} else if month == time.March {
		return 3
	} else if month == time.April {
		return 4
	} else if month == time.May {
		return 5
	} else if month == time.June {
		return 6
	} else if month == time.July {
		return 7
	} else if month == time.August {
		return 8
	} else if month == time.September {
		return 9
	} else if month == time.October {
		return 10
	} else if month == time.November {
		return 11
	} else if month == time.December {
		return 12
	} else {
		return -1
	}
}

func DateHelper_renderDate(date time.Time) string {
	monthNumber := DateHelper_monthAsNumber(date.Month())
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d", date.Year(), monthNumber, date.Day(), date.Hour(), date.Minute())
}

func DateHelper_renderDateWithSecond(date time.Time) string {
	monthNumber := DateHelper_monthAsNumber(date.Month())
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", date.Year(), monthNumber, date.Day(), date.Hour(), date.Minute(), date.Second())
}
