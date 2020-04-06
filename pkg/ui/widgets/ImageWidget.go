package widgets

import (
	"fmt"
	"net/http"
)

type ImageWidget struct {
	Source string /* Source */
	Alt    string
}

func (self *ImageWidget) SetSource(s string) {
	self.Source = s
}

func NewImageWidget() *ImageWidget {
	iw := new(ImageWidget)
	return iw
}

func (self *ImageWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<img src=\"%s\" alt=\"%s\" />\n", self.Source, self.Alt)
	return nil
}