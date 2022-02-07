package widgets

import (
	"fmt"
	"io"
	"strings"
)

type ActionMenuWidget struct {
	actions []*MenuAction /* Actions */
}

func (self *ActionMenuWidget) Add(action *MenuAction) *ActionMenuWidget {
	self.actions = append(self.actions, action)
	return self
}

func NewActionMenuWidget() *ActionMenuWidget {
	amw := new(ActionMenuWidget)
	return amw
}

func (self *ActionMenuWidget) Render(w io.Writer) error {

	fmt.Fprintf(w, "<div class=\"actions\">\n")

	for _, a := range self.actions {

		var ourClass []string
		ourClass = append(ourClass, "btn")
		ourClass = append(ourClass, a.Class)
		newClass := strings.Join(ourClass, " ")

		fmt.Fprintf(w, "\t<div class=\"action-cover\">\n")
		fmt.Fprintf(w, "\t\t<a href=\"%s\" class=\"%s\">\n", a.Link, newClass)
		fmt.Fprintf(w, "\t\t\t<span><i class=\"%s\"></i>%s</span>\n", a.Icon, a.Label)
		fmt.Fprintf(w, "\t\t</a>\n")
		fmt.Fprintf(w, "\t</div>\n")
	}

	fmt.Fprintf(w, "</div>\n")

	return nil
}
