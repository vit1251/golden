package widgets

import (
	"fmt"
	"net/http"
)

type DivWidget struct {
	Class   string
	Content string
	Widget  IWidget
}

func (self *DivWidget) SetClass(s string) *DivWidget {
	self.Class = s
	return self
}

func (self *DivWidget) SetWidget(w IWidget) *DivWidget {
	self.Widget = w
	return self
}

func (self *DivWidget) SetContent(s string) {
	self.Content = s
}

func NewDivWidget() *DivWidget {
	iw := new(DivWidget)
	return iw
}

func (self *DivWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<div class=\"%s\">", self.Class)
	if self.Widget != nil {
		self.Widget.Render(w)
	} else {
		fmt.Fprintf(w, "%s", self.Content)
	}
	fmt.Fprintf(w, "</div>\n")
	return nil
}
