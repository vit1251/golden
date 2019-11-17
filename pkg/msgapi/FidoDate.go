package msgapi

import (
	"log"
	"fmt"
)

type FidoDate struct {
	Date uint16   // DOS bitmapped date value
	Time uint16   // DOS bitmapped time  value
}

const (
    FidoDateDayMask   = 0x001F   // 00000000 00011111
    FidoDateMonthMask = 0x01E0   // 00000001 11100000
    FidoDateYearMask  = 0xFE00   // 11111110 00000000
)

/**
 * Get date
 *
 * The first five bits represent the day of the month.
 * (A value of 1 represents the first of the month.)
 *
 * The next four bits indicate the month of the year.
 * (1=January; 12=December.)
 *
 * The remaining seven bits indicate the year (relative to 1980).
 */
func (self FidoDate) GetDate() string {
	log.Printf("parseDate = %016b", self.Date)
	//
	day := self.Date & FidoDateDayMask
	log.Printf("day = %016b", day)
	log.Printf("day = %d", day)
	//
	month := self.Date & FidoDateMonthMask
	log.Printf("month = %016b", month)
	month = month >> 5
	log.Printf("month = %d", month)
	//
	year := self.Date & FidoDateYearMask
	log.Printf("year = %016b", year)
	year = year >> 9
	log.Printf("year = %d (i.e. %d)", year, year + 1980)
	year += 1980
	//
	result := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	//
	return result
}

const (
    FidoDateSecondsMask     = 0x001F   // 00000000 00011111
    FidoDateMinuteMask      = 0x07E0   // 00000111 11100000
    FidoDateHourMask        = 0xF800   // 11111000 00000000
)


/**
 * Get time
 *
 * The first five bits indicate the seconds
 * value,  divided by two. This  implies
 * that all message dates and times  get
 * rounded to a multiple of two seconds.
 * Example:
 *   0 seconds =  0
 *  16 seconds =  8
 *  58 seconds = 29
 *
 * The next six  bits represent the minutes value.
 *
 * The  remaining  five bits  represent the
 * hour value, using a 24-hour clock.
 */
func (self FidoDate) GetTime() string {
	//
	seconds := self.Time & FidoDateSecondsMask
	log.Printf("seconds = %016b", seconds)
	seconds = seconds << 1
	//
	minute := self.Time & FidoDateMinuteMask
	log.Printf("minute = %016b", minute)
	minute = minute >> 5
	//
	hour := self.Time & FidoDateHourMask
	log.Printf("hour = %016b", hour)
	hour = hour >> 11
	//
	result := fmt.Sprintf("%02d:%02d:%02d", hour, minute, seconds)
	return result
}

func (self FidoDate) GetDateTime() string {
	var res string = self.GetDate() + " " + self.GetTime()
	return res
}
