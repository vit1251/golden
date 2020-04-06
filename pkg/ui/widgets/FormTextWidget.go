package widgets

import (
	"fmt"
	"net/http"
)

type FormTextWidget struct {
	Name string
}

func (self *FormTextWidget) SetName(s string) *FormTextWidget {
	self.Name = s
	return self
}

func NewFormTextWidget() *FormTextWidget {
	ftw:= new(FormTextWidget)
	return ftw
}

func (self *FormTextWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<div>\n")
	fmt.Fprintf(w, "\t<textarea name=\"%s\"></textarea>\n", self.Name)
	fmt.Fprintf(w, "</div>\n")
	return nil
}

