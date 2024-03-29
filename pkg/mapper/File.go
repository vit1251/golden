package mapper

import (
	"strings"
	"time"
)

type File struct {
	area      string    /* Area name         */
	desc      string    /* Description       */
	file      string    /* Disk index        */
	orig_name string    /* Original filename */
	time      time.Time /* Creation stamp    */
	viewCount int       /* View count        */
	path      string    /* Absolute path     */
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

func (self *File) SetAbsolutePath(path string) {
	self.path = path
}

func (self *File) GetAbsolutePath() string {
	return self.path
}

func (self *File) SetOrigName(orig_name string) {
	self.orig_name = orig_name
}

func (self File) GetOrigName() string {
	return self.orig_name
}
