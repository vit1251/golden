package timezone

import (
	"log"
	"strconv"
	"time"
)

type TimeZoneParser struct {

}

func NewTimeZoneParser() *TimeZoneParser {
	tzp := new(TimeZoneParser)
	return tzp
}

func (self *TimeZoneParser) Parse(timezone string) (*time.Location, error) {

	if (timezone[0] != '-') && (timezone[0] != '+') {
		timezone = "+" + timezone
	}
	log.Printf("Parse zone: zone = %+v", timezone)

	//
	hours, err := strconv.Atoi(timezone[1:3])
	if err != nil {
		return nil, err
	}
	minutes, err := strconv.Atoi(timezone[4:])
	if err != nil {
		return nil, err
	}
	tzOffset := ((hours * 60) + minutes) * 60

	if timezone[0] == '-' {
		tzOffset = 0 - tzOffset
	}

	return time.FixedZone(timezone, tzOffset), nil
}
