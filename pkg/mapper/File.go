package mapper

import (
	"strings"
	"time"
)

type File struct {
	area      string    /* Area name        */
	desc      string    /* Description      */
	file      string    /* Filename         */
	time      time.Time /* Creation stamp   */
	viewCount int       /* View count       */
}

func NewFile() *File {
	newFile := new(File)
	newFile.time = time.Now()
	return newFile
}

func (self *File) SetArea(area string) {
	self.area = strings.ToUpper(area)
}

func (self *File) SetDesc(desc string) {
	self.desc = desc
}

func (self *File) SetFile(file string) {
	self.file = file
}

func (self *File) SetUnixTime(unixTime int64) {
	self.time = time.Unix(unixTime, 0)
}

func (self File) GetArea() string {
	return self.area
}

func (self File) GetFile() string {
	return self.file
}

func (self File) GetDesc() string {
	return self.desc
}

func (self File) GetUnixTime() int64 {
	return self.time.Unix()
}

func (self File) GetTime() time.Time {
	return self.time
}

func (self *File) SetViewCount(viewCount int) {
	self.viewCount = viewCount
}

func (self File) GetViewCount() int {
	return self.viewCount
}

func (self File) IsNew() bool {
	return self.viewCount == 0
}
