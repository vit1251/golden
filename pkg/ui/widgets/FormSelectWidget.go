package widgets

import (
	"fmt"
	"net/http"
)

type FormSelectWidget struct {
	name string
	options []*FormSelectOptionWidget
}

type FormSelectOptionWidget struct {
	name string
	value string
}

func (self *FormSelectWidget) SetName(name string) *FormSelectWidget {
	self.name = name
	return self
}

func (self *FormSelectOptionWidget) SetName(name string) *FormSelectOptionWidget {
	self.name = name
	return self
}

func (self *FormSelectOptionWidget) SetValue(value string) *FormSelectOptionWidget {
	self.value = value
	return self
}

func NewFormSelectWidget() *FormSelectWidget {
	s := new(FormSelectWidget)
	return s
}

func NewFormSelectOptionWidget() *FormSelectOptionWidget {
	o := new(FormSelectOptionWidget)
	return o
}

func (self *FormSelectWidget) AddOption(name string, value string) *FormSelectWidget {
	o := NewFormSelectOptionWidget().
		SetName(name).
		SetValue(value)
	self.options = append(self.options, o)
	return self
}

func (self *FormSelectWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<select name=\"%s\">\n", self.name)
	for _, o := range self.options {
		fmt.Fprintf(w, "\t<option value=\"%s\">%s</option>\n", o.value, o.name)
	}
	fmt.Fprintf(w, "</select>\n")
	return nil
}
