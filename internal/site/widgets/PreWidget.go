package widgets

import (
	"fmt"
	"html"
	"io"
)

type PreWidget struct {
	content   string
}

func (self *PreWidget) SetContent(content string) *PreWidget {
	self.content = content
	return self
}

func NewPreWidget() *PreWidget {
	iw := new(PreWidget)
	return iw
}

func (self *PreWidget) Render(w io.Writer) error {
	newContent := html.EscapeString(self.content)
	fmt.Fprintf(w, "<pre>%s</pre>", newContent)
	return nil
}
