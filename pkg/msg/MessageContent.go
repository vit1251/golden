package msg

import (
	"fmt"
	"github.com/vit1251/golden/pkg/packet"
)

type MessageContent struct {
	RAW    []byte
	charset  string
}

func NewMessageContent() (*MessageContent) {
	mc := new(MessageContent)
	return mc
}

func (self *MessageContent) AddLine(line string) {
	if (self.charset == "CP866") {
		newLine := fmt.Sprintf("%s\r\n", line)
		var rawLine []rune = []rune(newLine)
		chunk, err1 := packet.EncodeText(rawLine)
		if err1 != nil {
			panic(err1)
		}
		self.RAW = append(self.RAW, chunk...)
	} else {
		panic("Wrong charset")
	}
}

func (self *MessageContent) Pack() []byte {
	return self.RAW
}

func (self *MessageContent) SetCharset(charset string) {
	self.charset = charset
}
