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

/// Remove SEEN-BY message
func (self *MessageOriginTransformer) Transform(msg string) string {
	var result string
	rows := strings.Split(msg, "\r")
	for _, oneLine := range rows {
		if self.state == MessageStateBody {
			if strings.HasPrefix(oneLine, " * Origin: ") {
				self.state = MessageStateService
			} else {
				result += oneLine + "\r"
			}
		} else if self.state == MessageStateService {
			log.Printf("orgini: line = %+v", oneLine)
		} else {
			panic("unknown state")
		}
	}
	return result
}
