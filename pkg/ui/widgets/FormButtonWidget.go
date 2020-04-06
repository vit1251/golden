package widgets

import (
	"fmt"
	"net/http"
)

type FormButtonWidget struct {
	Type  string
	Title string
}

func NewFormButtonWidget() *FormButtonWidget {
	ftw:= new(FormButtonWidget)
	return ftw
}

func (self *FormButtonWidget) SetType(s string) *FormButtonWidget {
	self.Type = s
	return self
}

func (self *FormButtonWidget) SetTitle(s string) *FormButtonWidget {
	self.Title = s
	return self
}

func (self *FormButtonWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<div>\n")
	fmt.Fprintf(w, "\t<button class=\"btn\" type=\"%s\">%s</button>", self.Type, self.Title)
	fmt.Fprintf(w, "</div>\n")
	return nil
}

