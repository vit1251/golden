package fidotime

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

type TimeZoneParser struct {
	zoneReader *strings.Reader
}

func NewTimeZoneParser() *TimeZoneParser {
	tzp := new(TimeZoneParser)
	return tzp
}

func (self *TimeZoneParser) checkNext(next rune) (bool, error) {

	r, _, err := self.zoneReader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return false, err
		}
		panic(err)
	}

	if r == next {
		return true, nil
	}

	err2 := self.zoneReader.UnreadRune()

	return false, err2
}

func (self *TimeZoneParser) checkDigit() (*int, error) {

	r, _, err := self.zoneReader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		panic(err)
	}

	var value int

	if r == '0' {
		value = 0
		return &value, nil
	} else if r == '1' {
		value = 1
		return &value, nil
	} else if r == '2' {
		value = 2
		return &value, nil
	} else if r == '3' {
		value = 3
		return &value, nil
	} else if r == '4' {
		value = 4
		return &value, nil
	} else if r == '5' {
		value = 5
		return &value, nil
	} else if r == '6' {
		value = 6
		return &value, nil
	} else if r == '7' {
		value = 7
		return &value, nil
	} else if r == '8' {
		value = 8
		return &value, nil
	} else if r == '9' {
		value = 9
		return &value, nil
	}

	err2 := self.zoneReader.UnreadRune()

	return nil, err2
}

func (self *TimeZoneParser) checkComplete() (bool, error) {

	size := self.zoneReader.Size()
	pos, err1 := self.zoneReader.Seek(0, io.SeekCurrent)
	if err1 != nil {
		return false, err1
	}

	return size == pos, nil
}

// / Example: +0100, 0300, -0700
func (self *TimeZoneParser) Parse(zone string) (*time.Location, error) {

	log.Printf("TimeZoneParser: Parse: zone = %+v", zone)

	self.zoneReader = strings.NewReader(zone)

	var sign int = 1

	if ok, err := self.checkNext('-'); ok {
		sign = -1
	} else if err != nil {
		return nil, err
	}

	var hours int = 0

	if digit, err := self.checkDigit(); digit != nil && err == nil {
		hours = *digit
	} else {
		return nil, fmt.Errorf("timezone parse error")
	}
	if digit, err := self.checkDigit(); digit != nil && err == nil {
		hours = 10*hours + *digit
	} else {
		return nil, fmt.Errorf("timezone parse error")
	}

	var minutes int = 0

	if digit, err := self.checkDigit(); digit != nil && err == nil {
		minutes = *digit
	} else {
		return nil, fmt.Errorf("timezone parse error")
	}
	if digit, err := self.checkDigit(); digit != nil && err == nil {
		minutes = 10*minutes + *digit
	} else {
		return nil, fmt.Errorf("timezone parse error")
	}

	if ok, err := self.checkComplete(); ok && err == nil {
	} else {
		return nil, fmt.Errorf("timezone parse error")
	}

	self.zoneReader = nil

	tzOffset := sign * ((hours * 60) + minutes) * 60

	return time.FixedZone(zone, tzOffset), nil

}
