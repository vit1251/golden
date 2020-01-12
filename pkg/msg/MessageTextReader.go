package msg

import (
	"log"
	"unicode"
	"html/template"
	"strings"
)

type MessageState int

const MessageStateBody    = 1
const MessageStateService = 2

type MessageTextReader struct {
	State    MessageState
	result   string
}


func NewMessageTextReader() (*MessageTextReader) {
	mr := new(MessageTextReader)
	mr.State = MessageStateBody
	return mr
}

const (
	LineNoQuote        = 0
	LineQuoteLevel     = 1
	LineQuoteBody      = 2
)

func (self *MessageTextReader) ParseQuoteLine(quoteLine string) (author string, quoteLevel string, msg string) {

	msg = quoteLine

	return author, quoteLevel, msg

}

func (self *MessageTextReader) processLine(oneLine string) string {

	log.Printf("oneLine = %s", oneLine)

	if strings.HasPrefix(oneLine, " * Origin: ") {
		self.State = MessageStateService
	}

	var pureLine string
	var quoteAuthor string
	var startLine string
	var quoteMarkers string
	var quoteLine string
	var quoteLevel int = 0
	var quoteProbe bool = true

	var state = LineNoQuote
	for _, ch := range oneLine {

		pureLine += string(ch)

		if state == LineNoQuote {

			if quoteProbe && quoteAuthor != "" && ch == '>' {

				quoteLevel += 1
				state = LineQuoteLevel
				quoteMarkers += string(ch)

			} else if unicode.IsUpper(ch) {
				startLine += string(ch)
				quoteAuthor = quoteAuthor + string(ch)
			} else if unicode.IsSpace(ch) {
				startLine += string(ch)
			} else {
				startLine += string(ch)
				quoteProbe = false
			}

		} else if state == LineQuoteLevel {

			if ch == '>' {

				quoteLevel += 1
				quoteMarkers += string(ch)

			} else {

				quoteLine += string(ch)
				state = LineQuoteBody

			}

		} else if state == LineQuoteBody {

			quoteLine += string(ch)

		}
	}

	var newLine string 

	if quoteLevel == 0 {

		newLine = pureLine

	} else {

		if quoteLevel % 2 == 0 {
			newLine = "<span style='color: red'>" + startLine + quoteMarkers + quoteLine + "</span>"
		} else {
			newLine = "<span style='color: green'>" + startLine +  quoteMarkers + quoteLine + "</span>"
		}

	}

	return newLine
}

func (self *MessageTextReader) Prepare(msg string) template.HTML {

	/* Replace CRLF on \x0D */
	newMsg := msg
	newMsg = strings.ReplaceAll(newMsg, "\x0A\x0D", "\x0D")
	newMsg = strings.ReplaceAll(newMsg, "\x0A", "\x0D")

	/* Process */
	var oneLine string
	for _, ch := range newMsg {
		if ch == '\x0D' {

			if self.State == MessageStateBody {
				newLine := self.processLine(oneLine)
				self.result += string(newLine) + "<br>"
			} else {
				log.Printf("Service line: row = %s", oneLine)
			}

			oneLine = ""

		} else {
			oneLine = oneLine + string(ch)
		}
	}

	return template.HTML(self.result)
}