package widgets

import (
	"fmt"
	"io"
	"strings"
)

type DivWidget struct {
	className string
	styles    []string
	widgets   []IWidget
	title     string
}

func (self *DivWidget) SetClass(s string) *DivWidget {
	self.className = s
	return self
}

func (self *DivWidget) AddWidget(w IWidget) *DivWidget {
	self.widgets = append(self.widgets, w)
	return self
}

// / deprecate
func (self *DivWidget) SetContent(content string) *DivWidget {
	self.AddWidget(NewTextWidgetWithText(content))
	return self
}

func NewDivWidget() *DivWidget {
	iw := new(DivWidget)
	return iw
}

func (self *DivWidget) Render(w io.Writer) error {
	styles := strings.Join(self.styles, ";")
	fmt.Fprintf(w, "<div style=\"%s\" class=\"%s\">", styles, self.className)
	for _, widget := range self.widgets {
		widget.Render(w)
	}
	fmt.Fprintf(w, "</div>\n")
	return nil
}

func (self *DivWidget) SetHeight(height string) *DivWidget {
	row := fmt.Sprintf("height: %s", height)
	return self.SetStyle(row)
}

func (self *DivWidget) SetWidth(width string) *DivWidget {
	row := fmt.Sprintf("width: %s", width)
	return self.SetStyle(row)
}

func (self *DivWidget) SetStyle(row string) *DivWidget {
	self.styles = append(self.styles, row)
	return self
}

func (self *DivWidget) SetTitle(title string) *DivWidget {
	self.title = title
	return self
}
