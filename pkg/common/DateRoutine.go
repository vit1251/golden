package commonfunc

import (
	"fmt"
	"time"
)

func MakeHumanTime(newTime time.Time) string {

	nowTime := time.Now()

	result := "Now"

	if nowTime.Year() != newTime.Year() {
		result = fmt.Sprintf("%d-%d-%d %02d:%02d", newTime.Year(), newTime.Month(), newTime.Day(), newTime.Hour(), newTime.Minute())
	} else if nowTime.Month() != newTime.Month() {
		result = fmt.Sprintf("%d-%d %02d:%02d", newTime.Month(), newTime.Day(), newTime.Hour(), newTime.Minute())
	} else if nowTime.Day() != newTime.Day() {
		result = fmt.Sprintf("%d %02d:%02d", newTime.Day(), newTime.Hour(), newTime.Minute())
	} else {
		result = fmt.Sprintf("%02d:%02d", newTime.Hour(), newTime.Minute())
	}

	return result
}

