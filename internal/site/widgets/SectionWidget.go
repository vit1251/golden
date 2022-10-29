package widgets

import (
	"fmt"
	"io"
)

type SectionWidget struct {
	Title  string
	Widget IWidget
}

func (self *SectionWidget) SetTitle(s string) *SectionWidget {
	self.Title = s
	return self
}

func (self *SectionWidget) SetWidget(w IWidget) *SectionWidget {
	self.Widget = w
	return self
}

func NewSectionWidget() *SectionWidget {
	sw := new(SectionWidget)
	return sw
}

func (self *SectionWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<section>")
	fmt.Fprintf(w, "<h1>%s</h1>", self.Title)
	self.Widget.Render(w)
	fmt.Fprintf(w, "</section>")
	return nil
}
