package msg

import "strings"

type MessageContentParser struct {
}

func NewMessageContentParser() *MessageContentParser {
	return new(MessageContentParser)
}

const (
	CR = "\x0D"
	LF = "\x0A"
)

func (self MessageContentParser) Parse(content string) (*MessageContent, error) {

	mc := NewMessageContent(nil)

	rows := strings.Split(content, CR)
	for _, row := range rows {
		mc.AddLine(row)
	}

	return mc, nil

}
