package msg

import (
	"testing"
)

func TestParseQuoteLineMessage1(t *testing.T) {

	msgProc := NewMessageTextProcessor()

	var origMsg string = "Hello, All!"

	author, quoteLevel, msg :=  msgProc.ParseQuoteLine(origMsg)
	if author != "" {
	}
	if quoteLevel != "" {
	}
	if msg != origMsg {
	}

}

//func TestParseQuoteLineMessage2(t *testing.T) {
//	var row2 string = ""
//}

//func TestParseQuoteLineMessage3(t *testing.T) {
//	var row3 string = " VS> Level1"
//}

//func TestParseQuoteLineMessage4(t *testing.T) {
//	var row4 string = " VS>> Level2"
//}

//func TestParseQuoteLineMessage5(t *testing.T) {
//	var row5 string = "> Quote"
//}

func TestParseQuoteLineMessage6(t *testing.T) {

	msgProc := NewMessageTextProcessor()

	var origMsg string = "Line contain > in body"

	author, quoteLevel, msg :=  msgProc.ParseQuoteLine(origMsg)
	if author != "" {
	}
	if quoteLevel != "" {
	}
	if msg != origMsg {
	}

}

