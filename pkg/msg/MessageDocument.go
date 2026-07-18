package msg

import (
    "html/template"
)

type MessageDocumentElement struct {
    Author     string
    QuoteLevel int
    Text       string
}

type MessageDocument struct {
    items []MessageDocumentElement
}

func (self *MessageDocument) Add(element MessageDocumentElement) {
    self.items = append(self.items, element)
}

func (self *MessageDocument) HTML() template.HTML {
    r := NewMessageDocumentRenderer()
    return r.RenderAsHTML(self)
}

func (self *MessageDocument) Content() string {
    r := NewMessageDocumentRenderer()
    return r.RenderAsText(self)
}
