package msg

import (
	"fmt"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/registry"
)

type MessageContent struct {
	RAW      []byte
	charset  string
	registry *registry.Container
}

func (self *MessageManager) NewMessageContent(r *registry.Container) *MessageContent {
	mc := new(MessageContent)
	mc.registry = r
	return mc
}

func (self *MessageContent) AddLine(line string) {

	charsetManager := self.restoreCharsetManager()

	if self.charset == "CP866" {
		newLine := fmt.Sprintf("%s\r", line)
		var rawLine []rune = []rune(newLine)
		chunk, err1 := charsetManager.Encode(rawLine)
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

func (self *MessageContent) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}