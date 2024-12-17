package tracker

import (
	"fmt"
	"strings"
	"time"
)

type TicBuilder struct {
	lines []string
}

func NewTicBuilder() *TicBuilder {
	return new(TicBuilder)
}

func (self *TicBuilder) SetArea(areaName string) {
	/* Area RU.GOLDEN */
	self.addLine(fmt.Sprintf("Area %s", areaName))
}

func (self *TicBuilder) addLine(line string) {
	self.lines = append(self.lines, line)
}

func (self TicBuilder) Build() string {
	return strings.Join(self.lines, CRLF)
}

func (self *TicBuilder) SetOrigin(origin string) {
	/* Origin 2:5023/24.3752 */
	self.addLine(fmt.Sprintf("Origin %s", origin))
}

func (self *TicBuilder) SetFrom(from string) {
	/* From 2:5023/24.3752 */
	self.addLine(fmt.Sprintf("From %s", from))
}

func (self *TicBuilder) SetFile(filename string) {
	/* File GoldenPoint-20200423.zip */
	self.addLine(fmt.Sprintf("File %s", filename)) // TODO - make 8.3 name ...
	self.addLine(fmt.Sprintf("LFile %s", filename))
}

func (self *TicBuilder) SetSize(size int64) {
	/* Size 0 */
	self.addLine(fmt.Sprintf("Size %d", size))
}

func (self *TicBuilder) SetPw(passwd string) {
	/* Pw ****** */
	self.addLine(fmt.Sprintf("Pw %s", passwd))
}

func (self *TicBuilder) SetCrc(crc string) {
	/* Crc ****** */
	self.addLine(fmt.Sprintf("Crc %s", crc))
}

func (self *TicBuilder) SetTo(to string) {
	/* To ****** */
	self.addLine(fmt.Sprintf("To %s", to))
}

func (self *TicBuilder) SetDesc(desc string) {
	/* Desc Golden Point - Night - 2020-04-23 */
	self.addLine(fmt.Sprintf("Desc %s", desc))
}

// / SetLDesc add long description
func (self *TicBuilder) SetLDesc(ldesc string) {
	var CR string = "\x0D"
	var LF string = "\x0A"
	var rows []string
	if strings.Contains(ldesc, CR+LF) {
		/*  MS-DOS, OS/2, Microsoft Windows, Symbian OS and etc. */
		rows = strings.Split(ldesc, CR+LF)
	} else if strings.Contains(ldesc, LF) {
		/* GNU/Linux, AIX, Xenix, Mac OS X, FreeBSD and etc. */
		rows = strings.Split(ldesc, LF)
	} else {
		/* Single */
		rows = append(rows, ldesc)
	}
	for _, row := range rows {
		self.addLine(fmt.Sprintf("LDesc %s", row))
	}
}

func (self *TicBuilder) AddSeenby(addr string) {
	self.addLine(fmt.Sprintf("Seenby %s", addr))
}

func (self *TicBuilder) SetDate(time time.Time) {
	self.addLine(fmt.Sprintf("Date %d", time.Unix()))
}

func (self *TicBuilder) AddPath(path string) {
	self.addLine(fmt.Sprintf("Path %s", path))
}
