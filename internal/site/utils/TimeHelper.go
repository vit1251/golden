package utils

import "fmt"

func TimeHelper_renderDurationInMilli(duration uint64) string {
	ms := duration % 1000
	duration = duration / 1000
	sec := duration % 60
	duration = duration / 60
	min := duration % 60
	duration = duration / 60
	hour := duration % 24
	duration = duration / 24
	if duration > 0 {
		// TODO - there is days ...
	}

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hour, min, sec, ms)
}
