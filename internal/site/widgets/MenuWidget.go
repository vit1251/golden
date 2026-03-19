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
	fmt.Fprintf(w, "<div>")
	fmt.Fprintf(w, "<div class=\"Header\">")
	for _, g := range self.groups {
		fmt.Fprintf(w, "<div class=\"Header-item-group\">")
		for _, item := range g.actions {
			fmt.Fprintf(w, "<a class=\"nav-link\" href=\"%s\">", item.Link)
			fmt.Fprintf(w, "<div class=\"Header-item\">")
			if item.Icon != "" {
				fmt.Fprintf(w, "<svg viewBox=\"0 0 24 24\" width=\"24\" height=\"24\">")
				fmt.Fprintf(w, "  <use href=\"/static/sprite.svg#%s\"></use>", item.Icon)
				fmt.Fprintf(w, "</svg>")
			} else {
				fmt.Fprintf(w, "<span class=\"tab-label\">%s</span>", item.Label)
			}
			fmt.Fprintf(w, "</div>")
			fmt.Fprintf(w, "</a>")
		}
		fmt.Fprintf(w, "</div>")
	}
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "</div>")
	return nil
}
