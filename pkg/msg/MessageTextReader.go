package msg

import (
	"html/template"
	"log"
	"strings"
	"unicode"
)

// Message text is unbounded and null terminated (note exception below).
//
// A 'hard' carriage return, 0DH,  marks the end of a paragraph, and must
// be preserved.
//
// So   called  'soft'  carriage  returns,  8DH,  may  mark  a   previous
// processor's  automatic line wrap, and should be ignored.  Beware  that
// they may be followed by linefeeds, or may not.
//
// All  linefeeds, 0AH, should be ignored.  Systems which display message
// text should wrap long lines to suit their application.
//
// If the first character of a physical line (e.g. the first character of
// the  message text, or the character immediately after a hard  carriage
// return (ignoring any linefeeds)) is a ^A (<control-A>, 01H), then that
// line  is  not  displayed  as  it  contains  control  information.  The
// convention for such control lines is:
//   o They begin with ^A
//   o They end at the end of the physical line (i.e. ignore soft <cr>s).
//   o They begin with a keyword followed by a colon.
//   o The keywords are uniquely assigned to applications.
//   o They keyword/colon pair is followed by application specific data.
//
// Current ^A keyword assignments are:
//   o TOPT <pt no> - destination point address
//   o FMPT <pt no> - origin point address
//   o INTL <dest z:n/n> <orig z:n/n> - used for inter-zone address

type MessageState int

const MessageStateBody    = 1
const MessageStateService = 2

type MessageTextProcessor struct {
	State    MessageState
	result   string
}


func NewMessageTextProcessor() *MessageTextProcessor {
	mr := new(MessageTextProcessor)
	mr.State = MessageStateBody
	return mr
}

const (
	LineNoQuote        = 0
	LineQuoteLevel     = 1
	LineQuoteBody      = 2
)

func (self *MessageTextProcessor) ParseQuoteLine(quoteLine string) (author string, quoteLevel string, msg string) {

	msg = quoteLine

	return author, quoteLevel, msg

}

func (self *MessageTextProcessor) processLine(oneLine string) string {

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

func (self *MessageTextProcessor) Prepare(msg string) error {

	/* Replace CRLF on \x0D */
	newMsg := msg
	newMsg = strings.ReplaceAll(newMsg, "\x0A", "")
	newMsg = strings.ReplaceAll(newMsg, "\x07", "&#8226;") // Bullet char

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

	return nil
}

func (self *MessageTextProcessor) MakeReply() error {
	// TODO - process each line and make quote ...
	return nil
}

func (self *MessageTextProcessor) HTML() template.HTML {
	return template.HTML(self.result)
}

func (self *MessageTextProcessor) Content() string {
	return self.result
}
