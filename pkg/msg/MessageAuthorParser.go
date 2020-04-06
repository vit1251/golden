package msg

import "unicode"

type MessageAuthorParser struct {
}

func NewMessageAuthorParser() *MessageAuthorParser {
	cmap := new(MessageAuthorParser)
	return cmap
}

type MessageAuthor struct {
	FirstName string
	LastName string
	QuoteName string
}

func (self *MessageAuthorParser) Parse(author string) (*MessageAuthor, error) {
	ma := new(MessageAuthor)
	for _, ch := range author {
		if unicode.IsUpper(ch) {
			ma.QuoteName += string(ch)
		}
	}
	return ma, nil
}
