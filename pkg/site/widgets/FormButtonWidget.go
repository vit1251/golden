package widgets

import (
	"fmt"
	"io"
)

type FormButtonWidget struct {
	buttonType  string
	title       string
	value       string
	name        string
	class       string
}

func NewFormButtonWidget() *FormButtonWidget {
	newFormButtonWidget := new(FormButtonWidget)
	newFormButtonWidget.class = "btn"
	return newFormButtonWidget
}

func (self *FormButtonWidget) SetType(s string) *FormButtonWidget {
	self.buttonType = s
	return self
}

func (self *FormButtonWidget) SetTitle(s string) *FormButtonWidget {
	self.title = s
	return self
}

func (self *FormButtonWidget) Render(w io.Writer) error {
	//fmt.Fprintf(w, "<div>\n")
	fmt.Fprintf(w, "\t<button name=\"%s\" value=\"%s\" class=\"%s\" type=\"%s\">%s</button>",
		self.name,
		self.value,
		self.class,
		self.buttonType,
		self.title)
	//fmt.Fprintf(w, "</div>\n")
	return nil
}

func (self *FormButtonWidget) SetName(name string) *FormButtonWidget {
	self.name = name
	return self
}

func (self *FormButtonWidget) SetValue(value string) *FormButtonWidget {
	self.value = value
	return self
}

