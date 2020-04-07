package widgets

import (
	"fmt"
	"net/http"
)

type TextWidget struct {
	content string
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

func (self *TextWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "%s", self.content)
	return nil
}
