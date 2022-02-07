package widgets

import (
	"fmt"
	"io"
)

type MenuGroup struct {
	actions []*MenuAction
}

type MenuWidget struct {
	groups []*MenuGroup
}

func (self *MenuWidget) Add(manuGroup *MenuGroup) {
	self.groups = append(self.groups, manuGroup)
}

func (self *MenuGroup) Add(manuAction *MenuAction) {
	self.actions = append(self.actions, manuAction)
}

func NewMenuWidget() *MenuWidget {
	mw := new(MenuWidget)
	return mw
}

func (self *MenuWidget) Render(w io.Writer) error {

	fmt.Fprintf(w, "<div>\n")

	fmt.Fprintf(w, "<div class=\"Header\">\n")

	for _, g := range self.groups {
		fmt.Fprintf(w, "\t<div class=\"Header-item-group\">\n")
		for _, item := range g.actions {
			fmt.Fprintf(w, "\t\t<a class=\"nav-link\" href=\"%s\">\n", item.Link)
			fmt.Fprintf(w, "\t\t\t<div class=\"Header-item\">\n")
			fmt.Fprintf(w, "\t\t\t\t<span class=\"tab-label\">%s</span>\n", item.Label)
			if item.Metric > 0 {
				fmt.Fprintf(w, "\t\t\t\t<span class=\"badge\" id=\"%s\">%d</span>\n", item.ID, item.Metric)
			} else {
				fmt.Fprintf(w, "\t\t\t\t<span class=\"badge hidden\" id=\"%s\"></span>\n", item.ID)
			}
			fmt.Fprintf(w, "\t\t\t</div>\n")
			fmt.Fprintf(w, "\t\t</a>\n")
		}
		fmt.Fprintf(w, "\t</div>\n")
	}

	fmt.Fprintf(w, "</div>\n")

	fmt.Fprintf(w, "</div>\n")

	return nil
}
