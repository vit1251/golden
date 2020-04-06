package msg

import "log"

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
	log.Printf("reply_transform: msg = %+v", content)

	tr := NewTextReader()
	tr.Process(content, func(oneLine string) {
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

		row := nl.QuoteStart + nl.QuoteAuthor + nl.QuoteMarkers + nl.QuoteLine
		log.Printf("quote: row = %+v", row)

		newContent += row + "\n"
	})
	return newContent
}

func (self *MessageReplyTransformer) SetAuthor(author string) {
	self.author = author
}
