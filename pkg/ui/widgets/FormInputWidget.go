package widgets

import (
	"fmt"
	"net/http"
)

type FormInputWidget struct {
	Title       string
	Name        string
	Placeholder string
	Value       string
}

func NewFormInputWidget() *FormInputWidget {
	fi:= new(FormInputWidget)
	return fi
}

func (self *FormInputWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<div>\n")
	fmt.Fprintf(w, "\t<div>%s</div>\n", self.Title)
	fmt.Fprintf(w, "\t<div><input type=\"text\" value=\"%s\" name=\"%s\" placeholder=\"%s\" />\n", self.Value, self.Name, self.Placeholder)
	fmt.Fprintf(w, "</div>\n")
	return nil
}

func (self *FormInputWidget) SetPlaceholder(s string) *FormInputWidget {
	self.Placeholder = s
	return self
}

func (self *FormInputWidget) SetName(s string) *FormInputWidget {
	self.Name = s
	return self
}

func (self *FormInputWidget) SetTitle(s string) *FormInputWidget {
	self.Title = s
	return self
}

func (self *FormInputWidget) SetValue(value string) *FormInputWidget {
	self.Value = value
	return self
}
