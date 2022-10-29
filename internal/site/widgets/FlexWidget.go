package widgets

import (
	"fmt"
	"io"
	"strings"
)

type FlexWidget struct {
	className string
	widgets   []IWidget
	styles    []string
}

func NewFlexWidget() *FlexWidget {
	w := new(FlexWidget)
	return w
}

func (self *FlexWidget) SetClass(className string) *FlexWidget {
	self.className = className
	return self
}

func (self *FlexWidget) AddWidget(w IWidget) *FlexWidget {
	self.widgets = append(self.widgets, w)
	return self
}

func (self *FlexWidget) Render(w io.Writer) error {
	styles := strings.Join(self.styles, ";")
	fmt.Fprintf(w, "<div style=\"%s\" class=\"container-fluid %s\">", styles, self.className)
	if self.widgets != nil {
		for _, widget := range self.widgets {
			widget.Render(w)
		}
	}
	fmt.Fprintf(w, "</div>\n")
	return nil
}

func (self *FlexWidget) SetStyle(row string) *FlexWidget {
	self.styles = append(self.styles, row)
	return self
}
