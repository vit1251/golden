package msg

import (
	"html/template"
	"log"
	"strings"
)

type MessageTextProcessor struct {
	html  string
	raw   string
}


func NewMessageTextProcessor() *MessageTextProcessor {
	mr := new(MessageTextProcessor)
	return mr
}

func (self *MessageTextProcessor) ParseQuoteLine(quoteLine string) (author string, quoteLevel string, msg string) {

	msg = quoteLine

	return author, quoteLevel, msg

}

func (self *MessageTextProcessor) processHtmlLine(oneLine string) string {

	log.Printf("oneLine = %s", oneLine)

	mlp := NewMessageLineParser()
	ml := mlp.Parse(oneLine)

	var newLine string 

	if ml.QuoteLevel == 0 {
		newLine = ml.PureLine
	} else {
		if ml.QuoteLevel % 2 == 0 {
			newLine = "<span style='color: red'>" + ml.QuoteStart + ml.QuoteAuthor + ml.QuoteMarkers + ml.QuoteLine + "</span>"
		} else {
			newLine = "<span style='color: green'>" + ml.QuoteStart + ml.QuoteAuthor + ml.QuoteMarkers + ml.QuoteLine + "</span>"
		}
	}

	return newLine
}

func (self *MessageTextProcessor) Prepare(msg string) error {

	newMsg := msg
	newMsg = strings.ReplaceAll(newMsg, "\r\n", "\r")
	newMsg = strings.ReplaceAll(newMsg, "\x07", "&#8226;")

	/* Process */
	rows := strings.Split(msg, "\r")
	for _, oneLine := range rows {
		var newHtmlLine string = self.processHtmlLine(oneLine)
		var newLine = oneLine
		self.html += newHtmlLine + "<br>"
		self.raw += newLine + "\r"
	}

	return nil
}

func (self *MessageTextProcessor) HTML() template.HTML {
	return template.HTML(self.html)
}

func (self *MessageTextProcessor) Content() string {
	return self.raw
}
