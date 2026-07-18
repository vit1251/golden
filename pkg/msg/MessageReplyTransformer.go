package msg

import (
    "strings"
)

type MessageReplyTransformer struct {
    author string
}

func NewMessageReplyTransformer() *MessageReplyTransformer {
    mrt := new(MessageReplyTransformer)
    mrt.author = "??"
    return mrt
}

func (t *MessageReplyTransformer) SetAuthor(author string) { t.author = author }

func (t *MessageReplyTransformer) Transform(content string) string {
    var out string
    lines := StringSplitLines(content)
    for _, line := range lines {
        // Шаг 1. Парсим строки
    	mlp := NewMessageLineParser()
	nl := mlp.Parse(line)
        if nl.QuoteLevel == 0 {
	   nl.Author = t.author
        }
        nl.QuoteLevel += 1

        // Шаг 2. Собираем результат
        row := " " + nl.Author + strings.Repeat(">", nl.QuoteLevel+1) + nl.Text
        out = out + row + "\n"
    }

    return out
}
