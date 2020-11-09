package widgets

import (
	"fmt"
	"io"
)

type HeaderWidget struct {
	level int
	title string
}

func NewHeaderWidget() *HeaderWidget {
	w := new(HeaderWidget)
	w.level = 1
	return w
}

func (self *HeaderWidget) SetTitle(title string) *HeaderWidget {
	self.title = title
	return self
}

func (self *HeaderWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<h%d>%s</h%d>\n", self.level, self.title, self.level)
	return nil
}
