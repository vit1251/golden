package msg

import (
	"log"
	"strings"
)

type MessageTransformState int

type MessageOriginTransformer struct {
	state MessageTransformState
}

const (
	MessageStateBody     MessageTransformState = 1
	MessageStateService  MessageTransformState = 2
)

func NewMessageOriginTransformer() *MessageOriginTransformer {
	ot := new(MessageOriginTransformer)
	ot.state = MessageStateBody
	return ot
}

func (self *MessageOriginTransformer) Transform(msg string) string {
	var result string
	tr := NewTextReader()
	tr.Process(msg, func(oneLine string) {
		if self.state == MessageStateBody {
			if strings.HasPrefix(oneLine, " * Origin: ") {
				self.state = MessageStateService
			} else {
				result += oneLine + "\n"
			}
		} else if self.state == MessageStateService {
			log.Printf("orgini: line = %+v", oneLine)
		} else {
			panic("unknown state")
		}
	})
	return result
}
