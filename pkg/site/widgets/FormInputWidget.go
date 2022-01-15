package widgets

import (
	"fmt"
	"io"
)

type FormInputWidget struct {
	Title       string
	Name        string
	Placeholder string
	Value       string
	class       string
}

func NewFormInputWidget() *FormInputWidget {
	fi := new(FormInputWidget)
	return fi
}

func (self *FormInputWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<div class=\"input\">\n")
	fmt.Fprintf(w, "\t<div>%s</div>\n", self.Title)
	fmt.Fprintf(w, "\t<div>")
	fmt.Fprintf(w, "\t\t<input class=\"%s\" type=\"text\" value=\"%s\" name=\"%s\" placeholder=\"%s\" />\n", self.class, self.Value, self.Name, self.Placeholder)
	fmt.Fprintf(w, "\t</div>")
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

func (self *FormInputWidget) SetClass(class string) *FormInputWidget {
	self.class = class
	return self
}
