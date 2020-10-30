package file

import (
	"github.com/xeonx/timeago"
	"strings"
	"time"
)

type TicFile struct {
	From        string
	To          string
	File        string
	area        string
	Desc        string
	SeenBy      []string
	UnixTime    int64
	DateWritten *time.Time
}

func NewTicFile() *TicFile {
	tic := new(TicFile)
	return tic
}

func (self *TicFile) AddSeenBy(sb string) {
	self.SeenBy = append(self.SeenBy, sb)
}

func (self *TicFile) SetUnixTime(unixTime int64) {
	self.UnixTime = unixTime
	tm := time.Unix(unixTime, 0)
	self.DateWritten = &tm
}

func (self *TicFile) Age() string {
	var result string = "-"
	if self.DateWritten != nil {
		result = timeago.English.Format(*self.DateWritten)
	}
	return result
}

func (self *TicFile) SetArea(area string) {
	self.area = strings.ToUpper(area)
}

func (self *TicFile) GetArea() string {
	return self.area
}