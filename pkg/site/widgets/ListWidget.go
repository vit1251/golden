package widgets

import (
	"fmt"
	"net/http"
)

type ListWidget struct {
	Class   string
	Items   []IWidget
}

func NewListWidget() *ListWidget {
	iw := new(ListWidget)
	return iw
}

func (self *ListWidget) SetClass(s string) *ListWidget {
	self.Class = s
	return self
}

func (self *ListWidget) AddItem(item IWidget) *ListWidget {
	self.Items = append(self.Items, item)
	return self
}

func (self ListWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<ul class=\"%s\">", self.Class)
	for _, i := range self.Items {
		fmt.Fprintf(w, "\t<li>\n")
		i.Render(w)
		fmt.Fprintf(w, "\t</li>\n")
	}
	fmt.Fprintf(w, "</ul>\n")
	return nil
}
