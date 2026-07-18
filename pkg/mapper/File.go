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
    origin    string    /* Source address    */
    from      string    /* Addr sender       */
    to        string    /* Addr receiver     */
    size      int64     /* Size (by TIC)     */
    crc       string    /* CRC16 / CRC32     */
    ldesc     string    /* Long description  */
}

func NewFile() *File {
	newFile := new(File)
	newFile.time = time.Now()
	return newFile
}

func (self *File) SetArea(area string) { self.area = strings.ToUpper(area) }
func (self *File) SetDesc(desc string) { self.desc = desc }
func (self *File) SetFile(file string) { self.file = file }
func (self *File) SetUnixTime(unixTime int64) { self.time = time.Unix(unixTime, 0) }
func (self File) GetArea() string { return self.area }
func (self File) GetFile() string { return self.file }
func (self File) GetDesc() string { return self.desc }
func (self File) GetUnixTime() int64 { return self.time.Unix() }
func (self File) GetTime() time.Time { return self.time }
func (self *File) SetViewCount(viewCount int) { self.viewCount = viewCount }
func (self File) GetViewCount() int { return self.viewCount }
func (self File) IsNew() bool { return self.viewCount == 0 }
func (self *File) SetAbsolutePath(path string) { self.path = path }
func (self *File) GetAbsolutePath() string { return self.path }
func (self *File) SetOrigName(orig_name string) { self.orig_name = orig_name }

func (self File) GetOrigName() string { return self.orig_name }
func (self *File) SetOrigin(v string) { self.origin = v }
func (self File) GetOrigin() string   { return self.origin }
func (self *File) SetFrom(v string)   { self.from = v }
func (self File) GetFrom() string     { return self.from }
func (self *File) SetTo(v string)     { self.to = v }
func (self File) GetTo() string       { return self.to }
func (self *File) SetSize(v int64)    { self.size = v }
func (self File) GetSize() int64      { return self.size }
func (self *File) SetCrc(v string)    { self.crc = v }
func (self File) GetCrc() string      { return self.crc }
func (self *File) SetLDesc(v string)  { self.ldesc = v }
func (self File) GetLDesc() string    { return self.ldesc }
