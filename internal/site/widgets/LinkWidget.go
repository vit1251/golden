package widgets

import (
	"fmt"
	"io"
)

type LinkWidget struct {
	Link       string
	classNames string
	widgets    []IWidget
}

func NewLinkWidget() *LinkWidget {
	lw := new(LinkWidget)
	return lw
}

func (self *LinkWidget) SetContent(content string) *LinkWidget {
	return self.AddWidget(NewTextWidgetWithText(content))
}

func (self *LinkWidget) SetLink(link string) *LinkWidget {
	self.Link = link
	return self
}

func (self *LinkWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<a href=\"%s\" class=\"%s\">", self.Link, self.classNames)
	for _, widget := range self.widgets {
		widget.Render(w)
	}
	fmt.Fprintf(w, "</a>")
	return nil
}

func (self *LinkWidget) SetClass(classNames string) *LinkWidget {
	self.classNames = classNames
	return self
}

func (self *LinkWidget) AddWidget(widget IWidget) *LinkWidget {
	self.widgets = append(self.widgets, widget)
	return self
}
