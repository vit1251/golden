package msg

import (
	"log"
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

func (self *MessageReplyTransformer) Transform(content string) string {
	var newContent string

	// Привет, Vitold!
	//
	// Четверг 05 Ноября 2020 13:02:06, Vitold Sedyshev писал(а) к Jaroslav Bespalov:
	//
	// С наилучшими пожеланиями, Jaroslav.

	log.Printf("reply_transform: msg = %+v", content)

	rows := strings.Split(content, "\r")
	for _, oneLine := range rows {

		log.Printf("transform: row = %+v", oneLine)

		mlp := NewMessageLineParser()
		nl := mlp.Parse(oneLine)

		if nl.QuoteLevel == 0 {
			nl.QuoteLevel = 1
			nl.QuoteAuthor = self.author
			nl.QuoteMarkers = ">"
			nl.QuoteStart = " "
			nl.QuoteLine = " " + nl.PureLine
		} else {
			nl.QuoteStart = " "
			nl.QuoteMarkers += ">"
			nl.QuoteLevel += 1
		}

		log.Printf("ql: nl = %+v", nl)

		quoteLineSize := len(strings.TrimSpace(nl.QuoteLine))
		if quoteLineSize > 0 {
			row := nl.QuoteStart + nl.QuoteAuthor + nl.QuoteMarkers + nl.QuoteLine
			newContent += row + "\r"
			log.Printf("quote: row = %+v", row)
		} else {
			newContent += "\r"
		}

	}

	return newContent
}

func (self *MessageReplyTransformer) SetAuthor(author string) {
	self.author = author
}
