package file

import "time"

type TicFile struct {
	From        string
	To          string
	File        string
	Area        string
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
