package msg

import (
	"fmt"
	"github.com/vit1251/golden/pkg/charset"
)

type MessageContent struct {
	cm      *charset.CharsetManager
	RAW     []byte
	charset string
}

func (self *MessageManager) NewMessageContent() *MessageContent {
	mc := new(MessageContent)
	mc.cm = self.CharsetManager
	return mc
}

func (self *MessageContent) AddLine(line string) {

	if self.charset == "CP866" {
		newLine := fmt.Sprintf("%s\r", line)
		var rawLine []rune = []rune(newLine)
		chunk, err1 := self.cm.Encode(rawLine)
		if err1 != nil {
			panic(err1)
		}
		self.RAW = append(self.RAW, chunk...)
	} else {
		panic("wrong charset")
	}
}

func (self *MessageContent) Pack() []byte {
	return self.RAW
}

func (self *MessageContent) SetCharset(charset string) {
	self.charset = charset
}
