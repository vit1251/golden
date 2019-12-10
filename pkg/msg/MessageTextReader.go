package msg

import (
	"log"
	"unicode"
	"html/template"
)

type MessageTextReader struct {
	result string
}


func NewMessageTextReader() (*MessageTextReader) {
	mr := new(MessageTextReader)
	return mr
}

const (
	LineUnkown     = 0
	LineQuotes     = 1
	LineQuoteBody  = 2
)

func (self *MessageTextReader) processLine(oneLine string) string {

	log.Printf("oneLine = %s", oneLine)

	var startLine string
	var newLine string
	var quoteLine string
	var quoteLevel int = 0
	var state = LineUnkown
	for _, ch := range oneLine {
		if state == LineUnkown {
			if ch == '>' {
				quoteLevel += 1
				state = LineQuotes
				quoteLine += string(ch)
			} else if unicode.IsUpper(ch) || unicode.IsSpace(ch) {
				startLine = startLine + string(ch)
			} else {
				newLine = startLine  + string(ch)
				state = LineQuoteBody
			}
		} else if state == LineQuotes {
			if ch == '>' {
				quoteLevel += 1
				quoteLine += string(ch)
			} else {
				newLine = newLine + string(ch)
				state = LineQuoteBody
			}
		} else {
			newLine = newLine + string(ch)
		}
	}

	if quoteLevel > 0 {
		if quoteLevel % 2 == 0 {
			newLine = "<span style='color: red'>" + startLine + quoteLine + newLine + "</span>"
		} else {
			newLine = "<span style='color: green'>" + startLine +  quoteLine + newLine + "</span>"
		}
	}

	return newLine
}

func (self *MessageTextReader) Prepare(msg string) template.HTML {

	var oneLine string
	for _, ch := range msg {
		if ch == '\x0A' || ch == '\x0D' {
			newLine := self.processLine(oneLine)
			oneLine = ""
			self.result += string(newLine) + "<br>"
		} else {
			oneLine = oneLine + string(ch)
		}
	}

	return template.HTML(self.result)
}