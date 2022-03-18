package msg

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"strings"
)

type MessageDocumentElement interface {
}

type MessageDocumentQuote struct {
	MessageDocumentElement
	authorName string
	quoteDeep  int
	origText   string
}

type MessageDocumentText struct {
	MessageDocumentElement
	origText string
}

type MessageDocument struct {
	items []MessageDocumentElement
}

type MessageTextProcessor struct {
}

func NewMessageTextProcessor() *MessageTextProcessor {
	mr := new(MessageTextProcessor)
	return mr
}

func (self *MessageTextProcessor) ParseQuoteLine(quoteLine string) (author string, quoteLevel string, msg string) {

	msg = quoteLine

	return author, quoteLevel, msg

}

func (self *MessageTextProcessor) processHtmlLine(oneLine string) MessageDocumentElement {

	log.Printf("oneLine = %s", oneLine)

	mlp := NewMessageLineParser()
	ml := mlp.Parse(oneLine)

	if ml.QuoteLevel == 0 {
		result := MessageDocumentText{
			origText: ml.PureLine,
		}
		return result
	} else {
		result := MessageDocumentQuote{
			authorName: ml.QuoteAuthor,
			quoteDeep:  ml.QuoteLevel,
			origText:   ml.QuoteLine,
		}
		return result
	}

}

func (self *MessageTextProcessor) Prepare(msg string) (*MessageDocument, error) {

	var doc *MessageDocument = new(MessageDocument)

	newMsg := msg
	newMsg = strings.ReplaceAll(newMsg, "\r\n", "\r")
	newMsg = strings.ReplaceAll(newMsg, "\x07", "&#8226;")

	/* Process */
	rows := strings.Split(msg, "\r")
	for _, oneLine := range rows {
		newElement := self.processHtmlLine(oneLine)
		doc.Add(newElement)
	}

	return doc, nil
}

func (self *MessageDocument) HTML() template.HTML {
	r := NewMessageDocumentRenderer()
	return r.RenderAsHTML(self)
}

func (self *MessageDocument) Content() string {
	r := NewMessageDocumentRenderer()
	return r.RenderAsText(self)
}

func (self *MessageDocument) Add(element MessageDocumentElement) {
	self.items = append(self.items, element)
}

type MessageDocumentRenderer struct {
}

func NewMessageDocumentRenderer() *MessageDocumentRenderer {
	return new(MessageDocumentRenderer)
}

func (self *MessageDocumentRenderer) RenderAsText(doc *MessageDocument) string {

	var items []string
	for _, item := range doc.items {
		switch v := item.(type) {
		case MessageDocumentText:
			newLine := fmt.Sprintf("%s", v.origText)
			items = append(items, newLine)
		case MessageDocumentQuote:
			var padding string
			for i := 0; i < v.quoteDeep; i++ {
				padding = padding + ">"
			}
			newLine := fmt.Sprintf(" %s%s%s", v.authorName, padding, v.origText)
			items = append(items, newLine)
		}
	}

	return strings.Join(items, "\r")
}

func (self *MessageDocumentRenderer) RenderAsHTML(doc *MessageDocument) template.HTML {

	var items []string
	for _, item := range doc.items {
		switch v := item.(type) {
		case MessageDocumentText:
			newText := html.EscapeString(v.origText)
			newLine := fmt.Sprintf("%s", newText)
			items = append(items, newLine)
		case MessageDocumentQuote:
			var padding string
			for i := 0; i < v.quoteDeep; i++ {
				padding = padding + ">"
			}
			var newLine string
			newText := html.EscapeString(v.origText)
			if v.quoteDeep%2 == 0 {
				newLine = "<span style='color: red'>" + v.authorName + padding + newText + "</span>"
			} else {
				newLine = "<span style='color: green'>" + v.authorName + padding + newText + "</span>"
			}
			items = append(items, newLine)
		}
	}
	var out string = strings.Join(items, "<br>")
	return template.HTML(out)
}
