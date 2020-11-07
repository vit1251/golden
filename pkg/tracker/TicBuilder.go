package tracker

import (
	"fmt"
	"strings"
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

func (self *TicBuilder) SetDesc(desc string) {
	/* Desc Golden Point - Night - 2020-04-23 */
	self.addLine(fmt.Sprintf("Desc %s", desc))
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
