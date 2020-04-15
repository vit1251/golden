package widgets

import (
	"fmt"
	"log"
	"net/http"
)

type FormTextWidget struct {
	name  string
	value string
	class string
}

func (self *FormTextWidget) SetName(s string) *FormTextWidget {
	self.name = s
	return self
}

func NewFormTextWidget() *FormTextWidget {
	ftw := new(FormTextWidget)
	return ftw
}

func (self *FormTextWidget) SetClass(class string) *FormTextWidget {
	self.class = class
	return self
}

func (self *FormTextWidget) Render(w http.ResponseWriter) error {

	log.Printf("class = %+v", self.class)

	fmt.Fprintf(w, "<div>\n")
	fmt.Fprintf(w, "\t<textarea class=\"%s\" name=\"%s\">%s</textarea>\n", self.class, self.name, self.value)
	fmt.Fprintf(w, "</div>\n")

	return nil
}

func (self *FormTextWidget) SetValue(value string) *FormTextWidget {
	self.value = value
	return self
}
