package widgets

import (
	"fmt"
	"net/http"
)

type FormTextWidget struct {
	Name  string
	value string
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
	fmt.Fprintf(w, "\t<textarea name=\"%s\">%s</textarea>\n", self.Name, self.value)
	fmt.Fprintf(w, "</div>\n")
	return nil
}

func (self *FormTextWidget) SetValue(content2 string) *FormTextWidget {
	self.value = content2
	return self
}

