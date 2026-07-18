package msg

import (
    "strings"
    "unicode"
)

type MessageLineParser struct {
}

func NewMessageLineParser() *MessageLineParser {
    return &MessageLineParser{}
}

func (self *MessageLineParser) Parse(data string) *MessageLine {
    runes := []rune(data)
    pos := 0

    for pos < len(runes) && runes[pos] == ' ' {
        pos++
    }

    var author strings.Builder
    for pos < len(runes) && unicode.IsUpper(runes[pos]) {
        author.WriteRune(runes[pos])
        pos++
    }

    quoteLevel := 0
    for pos < len(runes) && runes[pos] == '>' {
        quoteLevel++
        pos++
    }

    if quoteLevel > 0 {
        return &MessageLine{
            Author:     author.String(),
            Text:       string(runes[pos:]),
            QuoteLevel: quoteLevel,
        }
    }
    return &MessageLine{Text: data}
}
