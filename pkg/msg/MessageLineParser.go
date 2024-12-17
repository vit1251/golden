package msg

import "unicode"

type LineParserState int

type MessageLineParser struct {
	State LineParserState
}

const (
	LineQuoteStart  LineParserState = 0
	LineQuoteAuthor LineParserState = 1
	LineQuoteLevel  LineParserState = 2
	LineQuoteBody   LineParserState = 3
)

func NewMessageLineParser() *MessageLineParser {
	mlp := new(MessageLineParser)
	return mlp
}

func (self *MessageLineParser) Parse(oneLine string) *MessageLine {

	var ml MessageLine
	var state LineParserState = LineQuoteStart
	//var probe bool = true

	for _, ch := range oneLine {

		ml.PureLine += string(ch)

		if state == LineQuoteStart {
			if unicode.IsSpace(ch) {
				ml.QuoteStart += string(ch)
			} else if unicode.IsUpper(ch) {
				ml.QuoteAuthor += string(ch)
			} else if ch == '>' {
				state = LineQuoteLevel
				ml.QuoteMarkers += string(ch)
				ml.QuoteLevel += 1
			} else {
				ml.QuoteStart = ""
				ml.QuoteAuthor = ""
				state = LineQuoteBody
				ml.QuoteLine = ml.PureLine
			}
		} else if state == LineQuoteLevel {
			if ch == '>' {
				ml.QuoteMarkers += string(ch)
				ml.QuoteLevel += 1
			} else {
				ml.QuoteLine += string(ch)
				state = LineQuoteBody
			}
		} else if state == LineQuoteBody {
			ml.QuoteLine += string(ch)
		}

	}

	return &ml
}
