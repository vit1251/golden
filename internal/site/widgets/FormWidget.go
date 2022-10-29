package widgets

import (
	"fmt"
	"io"
)

type FormWidget struct {
	Action  string
	Method  string
	Widget  IWidget
	enctype string
}

func (self *FormWidget) SetAction(s string) *FormWidget {
	self.Action = s
	return self
}

func (self *FormWidget) SetMethod(s string) *FormWidget {
	self.Method = s
	return self
}

func (self *FormWidget) SetWidget(w IWidget) *FormWidget {
	self.Widget = w
	return self
}

func NewFormWidget() *FormWidget {
	fw:= new(FormWidget)
	fw.enctype = "application/x-www-form-urlencoded"
	return fw
}

func (self *FormWidget) Render(w io.Writer) error {
	fmt.Fprintf(w, "<form action=\"%s\" method=\"%s\" enctype=\"%s\">\n", self.Action, self.Method, self.enctype)
	if self.Widget != nil {
		self.Widget.Render(w)
	}
	fmt.Fprintf(w, "</form>\n")
	return nil
}

func (self *FormWidget) SetEnctype(enctype string) *FormWidget {
	self.enctype = enctype
	return self
}
