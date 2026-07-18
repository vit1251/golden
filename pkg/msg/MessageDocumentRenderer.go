package msg

import (
    "html"
    "html/template"
    "strings"
)

type MessageDocumentRenderer struct {
}

func NewMessageDocumentRenderer() *MessageDocumentRenderer {
	return new(MessageDocumentRenderer)
}

func PadEnd(s string, level int) string {
    padding := s
    for i := 0; i < level; i++ {
	padding = padding + ">"
    }
    return padding
}

func (self *MessageDocumentRenderer) RenderAsText(doc *MessageDocument) string {
    var lines []string
    for _, item := range doc.items {

        //
        if item.QuoteLevel == 0 {
	    lines = append(lines, item.Text)
	}

        //
        if item.QuoteLevel > 0 {
            if item.Author != "" {
                quoteBlock := item.Author + strings.Repeat(">", item.QuoteLevel)
                newLine := quoteBlock + " " + item.Text
                lines = append(lines, newLine)
            } else {
                quoteBlock := strings.Repeat(">", item.QuoteLevel)
                newLine := " " + quoteBlock + " " + item.Text
                lines = append(lines, newLine)
            }
        }

    }
    return strings.Join(lines, "\n")
}


func (self *MessageDocumentRenderer) renderLine(str string, level int) string {
    var newLine string
    if level%2 == 0 {
	newLine = "<span style='color: red'>" + str + "</span>"
    } else {
	newLine = "<span style='color: green'>" + str + "</span>"
    }
    return newLine
}

func (self *MessageDocumentRenderer) RenderAsHTML(doc *MessageDocument) template.HTML {

    var lines []string
    for _, item := range doc.items {

        if item.QuoteLevel == 0 {
            newText := html.EscapeString(item.Text)
	    lines = append(lines, newText)
	}

        if item.QuoteLevel > 0 {
            if item.Author != "" {
                quoteBlock := item.Author + strings.Repeat(">", item.QuoteLevel)
                newLine := " " + quoteBlock + " " + item.Text
                newText := html.EscapeString(newLine)
                newHTML := self.renderLine(newText, item.QuoteLevel)
                lines = append(lines, newHTML)
            } else {
                quoteBlock := strings.Repeat(">", item.QuoteLevel)
                newLine := " " + quoteBlock + " " + item.Text
                newText := html.EscapeString(newLine)
                newHTML := self.renderLine(newText, item.QuoteLevel)
                lines = append(lines, newHTML)
            }
        }

    }
    var out string = strings.Join(lines, "<br>")
    return template.HTML(out)
}
