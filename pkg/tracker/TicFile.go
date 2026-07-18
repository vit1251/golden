package tracker

import (
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
    size        int64
    crc         string
    ldesc       string
}

func NewTicFile() *TicFile {
    newTicFile := new(TicFile)
    newTicFile.dateWritten = time.Now()
    return newTicFile
}

func (self *TicFile) SetArea(area string) { self.area = strings.ToUpper(area) }
func (self *TicFile) GetArea() string { return self.area }
func (self *TicFile) SetPw(pw string) { self.pw = pw }
func (self *TicFile) GetUnixTime() int64 { return self.dateWritten.Unix() }
func (self *TicFile) SetUnixTime(unixTime int64) { self.dateWritten = time.Unix(unixTime, 0) }
func (self *TicFile) SetDesc(desc string) { self.desc = desc }
func (self *TicFile) SetFile(file string) { self.file = file }
func (self *TicFile) SetLFile(value string) { self.lfile = value }
func (self *TicFile) GetFile() string { return self.file }
func (self *TicFile) GetLFile() string { return self.lfile }
func (self *TicFile) GetDesc() string { return self.desc }
func (self *TicFile) GetOrigin() string { return self.origin }
func (self *TicFile) SetOrigin(v string) { self.origin = v }
func (self *TicFile) GetFrom() string    { return self.from }
func (self *TicFile) SetFrom(v string)   { self.from = v }
func (self *TicFile) GetTo() string      { return self.to }
func (self *TicFile) SetTo(v string)     { self.to = v }
func (self *TicFile) GetSize() int64     { return self.size }
func (self *TicFile) SetSize(v int64)    { self.size = v }
func (self *TicFile) GetCrc() string     { return self.crc }
func (self *TicFile) SetCrc(v string)    { self.crc = v }
func (self *TicFile) GetLDesc() string   { return self.ldesc }
func (self *TicFile) SetLDesc(v string)  { self.ldesc = v }
