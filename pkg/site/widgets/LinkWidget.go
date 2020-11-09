package widgets

import (
	"fmt"
	"io"
)

type LinkWidget struct {
	Link    string
	Content string
	Widget  IWidget
	class   string
}

func NewLinkWidget() *LinkWidget {
	lw := new(LinkWidget)
	return lw
}

func (self *LinkWidget) SetContent(content string) *LinkWidget {
	self.Content = content
	return self
}

func (self *LinkWidget) SetLink(link string) *LinkWidget {
	self.Link = link
	return self
}

func (self *LinkWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<a href=\"%s\" class=\"%s\">", self.Link, self.class)
	if self.Widget != nil {
		self.Widget.Render(w)
	} else {
		fmt.Fprintf(w, "%s", self.Content)
	}
	fmt.Fprintf(w, "</a>")
	return nil
}

func (self *LinkWidget) SetClass(s string) *LinkWidget {
	self.class = s
	return self
}
