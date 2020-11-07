package tracker

import (
	commonfunc "github.com/vit1251/golden/pkg/common"
	"strings"
	"time"
)

type TicFile struct {
	origin      string
	from        string
	to          string
	lfile       string
	file        string
	area        string
	desc        string
	dateWritten time.Time
	pw          string
}

func NewTicFile() *TicFile {
	newTicFile := new(TicFile)
	newTicFile.dateWritten = time.Now()
	return newTicFile
}

func (self *TicFile) GetAge() string {
	result := commonfunc.MakeHumanTime(self.dateWritten)
	return result
}

func (self *TicFile) SetArea(area string) {
	self.area = strings.ToUpper(area)
}

func (self *TicFile) GetArea() string {
	return self.area
}

func (self *TicFile) SetPw(pw string) {
	self.pw = pw
}

func (self *TicFile) GetUnixTime() int64 {
	return self.dateWritten.Unix()
}

func (self *TicFile) SetUnixTime(unixTime int64) {
	self.dateWritten = time.Unix(unixTime, 0)
}

func (self *TicFile) SetDesc(desc string) {
	self.desc = desc
}

func (self *TicFile) SetFile(file string) {
	self.file = file
}

func (self *TicFile) SetLFile(value string) {
	self.lfile = value
}

func (self *TicFile) GetFile() string {
	return self.file
}

func (self *TicFile) GetLFile() string {
	if self.lfile != "" {
		return self.lfile
	}
	return self.file
}

func (self *TicFile) GetDesc() string {
	return self.desc
}
