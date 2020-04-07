package widgets

import (
	"fmt"
	"net/http"
)

type LinkWidget struct {
	Link string
	Content string
	Widget IWidget
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

func (self *LinkWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<a href=\"%s\">", self.Link)
	if self.Widget != nil {
		self.Widget.Render(w)
	} else {
		fmt.Fprintf(w, "%s", self.Content)
	}
	fmt.Fprintf(w, "</a>")
	return nil
}
