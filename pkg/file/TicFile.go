package file

import (
	"bufio"
	"fmt"
	commonfunc "github.com/vit1251/golden/pkg/common"
	"os"
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
	Pw           string
	lines        []string
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

func (self *TicFile) SetPw(passwd string) {

}

func (self *TicFile) AddLine(line string) {
	self.lines = append(self.lines, line)
}

func (self TicFile) Save(path string) error {
	stream, err1 := os.Create(path)
	if err1 != nil {
		return err1
	}

	cacheStream := bufio.NewWriter(stream)

	defer func() {
		cacheStream.Flush()
		stream.Close()
	}()

	for _, line := range self.lines {
		newLine := fmt.Sprintf("%s\r", line)
		cacheStream.Write([]byte(newLine))
	}

	return nil
}