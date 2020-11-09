package widgets

import (
	"fmt"
	"io"
)

type FormFileInputWidget struct {
	Title       string
	Name        string
	Placeholder string
	Value       string
	class       string
}

func NewFormFileInputWidget() *FormFileInputWidget {
	fi:= new(FormFileInputWidget)
	return fi
}

func (self *FormFileInputWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<div>\n")
	fmt.Fprintf(w, "\t<div>%s</div>\n", self.Title)
	fmt.Fprintf(w, "\t<div><input class=\"%s\" type=\"file\" value=\"%s\" name=\"%s\" placeholder=\"%s\" />\n", self.class, self.Value, self.Name, self.Placeholder)
	fmt.Fprintf(w, "</div>\n")
	return nil
}

func (self *FormFileInputWidget) SetPlaceholder(s string) *FormFileInputWidget {
	self.Placeholder = s
	return self
}

func (self *FormFileInputWidget) SetName(s string) *FormFileInputWidget {
	self.Name = s
	return self
}

func (self *FormFileInputWidget) SetTitle(s string) *FormFileInputWidget {
	self.Title = s
	return self
}

func (self *FormFileInputWidget) SetValue(value string) *FormFileInputWidget {
	self.Value = value
	return self
}

func (self *FormFileInputWidget) SetClass(class string) *FormFileInputWidget {
	self.class = class
	return self
}

