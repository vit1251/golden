package msg

import (
	"fmt"
)

type MessageContent struct {
	RAW    []byte
}

func NewMessageContent() (*MessageContent) {
	mc := new(MessageContent)
	return mc
}

func (self *MessageContent) AddLine(line string) {
	newLine := fmt.Sprintf("%s\r\n", line)
	var rawLine []byte = []byte(newLine)
	self.RAW = append(self.RAW, rawLine...)
}

func (self *MessageContent) Pack() ([]byte) {
	return self.RAW
}
