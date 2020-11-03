package file

import (
	commonfunc "github.com/vit1251/golden/pkg/common"
	"strings"
	"time"
)

type TicFile struct {
	From         string
	To           string
	File         string
	area         string
	Desc         string
	SeenBy       []string
	UnixTime     int64
	DateWritten  time.Time
}

func NewTicFile() *TicFile {
	tic := new(TicFile)
	return tic
}

func (self *TicFile) AddSeenBy(sb string) {
	self.SeenBy = append(self.SeenBy, sb)
}

func (self *TicFile) SetUnixTime(unixTime int64) {
	self.DateWritten = time.Unix(unixTime, 0)
}

func (self *TicFile) GetAge() string {
	result := commonfunc.MakeHumanTime(self.DateWritten)
	return result
}

func (self *TicFile) SetArea(area string) {
	self.area = strings.ToUpper(area)
}

func (self *TicFile) GetArea() string {
	return self.area
}