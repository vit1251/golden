package widgets

import (
	"fmt"
	"io"
	"strings"
)

type FormInputWidget struct {
	Title       string
	Name        string
	Placeholder string
	Value       string
	class       string
	disable     bool
}

func NewFormInputWidget() *FormInputWidget {
	fi := new(FormInputWidget)
	return fi
}

func (self *FormInputWidget) Render(w io.Writer) error {

	var out strings.Builder

	/* Input wrapper start */
	out.WriteString("<div class=\"input\">\n")

	/* Title */
	out.WriteString("<div>")
	out.WriteString(self.Title)
	out.WriteString("</div>")

	/* Input */
	out.WriteString("<div>")
	out.WriteString("<")
	out.WriteString("input ")
	out.WriteString(fmt.Sprintf(" type=\"%s\"", "text"))
	out.WriteString(fmt.Sprintf(" class=\"%s\"", self.class))
	out.WriteString(fmt.Sprintf(" name=\"%s\"", self.Name))
	out.WriteString(fmt.Sprintf(" value=\"%s\"", self.Value))
	out.WriteString(fmt.Sprintf(" placeholder=\"%s\"", self.Placeholder))
	if self.disable {
		out.WriteString(fmt.Sprintf(" disabled"))
	}
	out.WriteString(" />")
	out.WriteString("</div>")

	/* Input wrapper stop */
	out.WriteString("</div>")

	/* Write out */
	_, err := fmt.Fprintf(w, "%s", out.String())
	return err
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

func (self *FormInputWidget) SetDisable(yesno bool) {
	self.disable = yesno
}
