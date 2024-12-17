package widgets

import (
	"fmt"
	"io"
)

type TextWidget struct {
	content string
	class   string
}

func NewTextWidget() *TextWidget {
	tw := new(TextWidget)
	return tw
}

func NewTextWidgetWithText(content string) *TextWidget {
	tw := NewTextWidget()
	tw.content = content
	return tw
}

func (self *TextWidget) SetClass(class string) *TextWidget {
	self.class = class
	return self
}

func (self *TextWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<span class=\"%s\">%s</span>", self.class, self.content)
	return nil
}
