package msg

import (
	"html/template"
	"log"
	"strings"
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
