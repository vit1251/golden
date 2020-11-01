package fidotime

import (
	"time"
	"fmt"
	"bytes"
	"errors"
	"io"
)

type FidoDate struct {
	year     int
	month    time.Month
	day      int
	hour     int
	minute   int
	sec      int
	msec     int
//	tz
}

func NewFidoDate() (*FidoDate) {
	fd := new(FidoDate)
	return fd
}

func (self *FidoDate) SetNow() {
	//
	newTime := time.Now()
	//
	self.year = newTime.Year()
	self.month = newTime.Month()
	self.day = newTime.Day()
	//
	self.hour = newTime.Hour()
	self.minute = newTime.Minute()
	self.sec = newTime.Second()
	//
}

func (self FidoDate) FTSC() ([]byte) {
	// []byte("03 Jan 20  23:51:10\x00")
	var result []byte
	newMonth := self.month.String()
	newMonth = newMonth[0:3]
	newDate := fmt.Sprintf("%02d %s %02d  %02d:%02d:%02d", self.day, newMonth, self.year - 2000, self.hour, self.minute, self.sec)
	result = []byte(newDate)
	return result
}

func (self FidoDate) CreateTime(zone *time.Location) (*time.Time, error) {
	result := time.Date(self.year, self.month, self.day, self.hour, self.minute, self.sec, 0, zone)
	return &result, nil
}

type DateParser struct {
	stream   *bytes.Buffer
	date      FidoDate
}

func NewDateParser() *DateParser {
	result := new(DateParser)
	return result
}

func (self *DateParser) parseSpace() (error) {
	value, err1 := self.stream.ReadByte()
	if err1 != nil {
		panic(err1)
	}
	if value == ' ' {
		/* Parse space complete */
	} else {
		return errors.New("unable parse space")
	}
	return nil
}

func (self *DateParser) parseChar(ch byte) (error) {
	value, err1 := self.stream.ReadByte()
	if err1 != nil {
		panic(err1)
	}
	if value == ch {
		/* Parse char complete */
	} else {
		return errors.New("unable parse char")
	}
	return nil
}

func (self *DateParser) parseNumber() (int, error) {

	var result = 0

	for {
		value, err1 := self.stream.ReadByte()
		if err1 != nil {
			if err1 == io.EOF {
				break
			} else {
				return result, err1
			}
		}
//		log.Printf("byte = %c", value)
		if value == '0' {
			result = result * 10
			result = result + 0
		} else if value == '1' {
			result = result * 10
			result = result + 1
		} else if value == '2' {
			result = result * 10
			result = result + 2
		} else if value == '3' {
			result = result * 10
			result = result + 3
		} else if value == '4' {
			result = result * 10
			result = result + 4
		} else if value == '5' {
			result = result * 10
			result = result + 5
		} else if value == '6' {
			result = result * 10
			result = result + 6
		} else if value == '7' {
			result = result * 10
			result = result + 7
		} else if value == '8' {
			result = result * 10
			result = result + 8
		} else if value == '9' {
			result = result * 10
			result = result + 9
		} else {
			err2 := self.stream.UnreadByte()
			if err2 != nil {
				panic(err2)
			}
			break
		}
	}

//	log.Printf("number = %d", result)

	return result, nil
}

func isLetter(c byte) bool {
    return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func (self *DateParser) parseString() ([]byte, error) {

	var cache []byte

	for {
		value, err1 := self.stream.ReadByte()
		if err1 != nil {
			if err1 == io.EOF {
				break
			} else {
				return cache, err1
			}
		}
//		log.Printf("byte = %c", value)
		if isLetter(value) {
			cache = append(cache, value)
		} else {
			err2 := self.stream.UnreadByte()
			if err2 != nil {
				return cache, err2
			}
			break
		}
	}

//	log.Printf("string = %s", cache)

	return cache, nil

}

func (self *DateParser) parseMonth() (*time.Month, error) {
	var result time.Month
	if monthName, err := self.parseString(); err != nil {
		return nil, err
	} else {
		if bytes.Equal(monthName, []byte("Jan")) {
			result = time.January
		} else if bytes.Equal(monthName, []byte("Feb")) {
			result = time.February
		} else if bytes.Equal(monthName, []byte("Mar")) {
			result = time.March
		} else if bytes.Equal(monthName, []byte("Apr")) {
			result = time.April
		} else if bytes.Equal(monthName, []byte("May")) {
			result = time.May
		} else if bytes.Equal(monthName, []byte("Jun")) {
			result = time.June
		} else if bytes.Equal(monthName, []byte("Jul")) {
			result = time.July
		} else if bytes.Equal(monthName, []byte("Aug")) {
			result = time.August
		} else if bytes.Equal(monthName, []byte("Sep")) {
			result = time.September
		} else if bytes.Equal(monthName, []byte("Oct")) {
			result = time.October
		} else if bytes.Equal(monthName, []byte("Nov")) {
			result = time.November
		} else if bytes.Equal(monthName, []byte("Dec")) {
			result = time.December
		} else {
			return nil, errors.New("Unable determin month name")
		}
	}
	return &result, nil
}

func (self *DateParser) Parse(date []byte) (*FidoDate, error) {

	self.stream = bytes.NewBuffer(date)

	/* Parse example: 01 Dec 19  09:03:20 */

	/* Parse date */
	if value, err := self.parseNumber() ; err != nil {
		return nil, errors.New("Unable parse day")
	} else {
		self.date.day = value
	}
	if err := self.parseSpace(); err != nil {
		return nil, errors.New("Unable parse space after day")
	}
	if value, err := self.parseMonth() ; err != nil {
		return nil, errors.New("Unable parse month")
	} else {
		self.date.month = *value
	}
	if err := self.parseSpace(); err != nil {
		return nil, errors.New("Unable parse space after month")
	}
	if value, err := self.parseNumber() ; err != nil {
		return nil, errors.New("Unable parse year")
	} else {
		self.date.year = 2000 + value
	}

	/* Parse separator */
	if err := self.parseSpace(); err != nil {
		return nil, errors.New("Unable parse space between date and time")
	}
	if err := self.parseSpace(); err != nil {
		return nil, errors.New("Unable parse space between date and time")
	}

	/* Parse time */
	if value, err := self.parseNumber() ; err != nil {
		return nil, errors.New("Unable parse hour")
	} else {
		self.date.hour = value
	}
	if err := self.parseChar(':'); err != nil {
		return nil, errors.New("Unable parse space time separator")
	}
	if value, err := self.parseNumber() ; err != nil {
		return nil, errors.New("Unable parse minute")
	} else {
		self.date.minute = value
	}
	if err := self.parseChar(':'); err != nil {
		return nil, errors.New("Unable parse space time separator")
	}
	if value, err := self.parseNumber() ; err != nil {
		return nil, errors.New("Unable parse seconds")
	} else {
		self.date.sec = value
	}

	return &self.date, nil

}
