package widgets

import (
	"fmt"
	"io"
)

type ImageWidget struct {
	source string
	alt    string
	class  string
}

func (self *ImageWidget) SetSource(s string) *ImageWidget {
	self.source = s
	return self
}

func NewImageWidget() *ImageWidget {
	iw := new(ImageWidget)
	return iw
}

func (self *ImageWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<img src=\"%s\" alt=\"%s\" class=\"%s\" />\n", self.source, self.alt, self.class)
	return nil
}

func (self *ImageWidget) SetClass(class string) *ImageWidget {
	self.class = class
	return self
}
