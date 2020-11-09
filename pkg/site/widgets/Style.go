package widgets

import (
	"fmt"
	"strings"
)

type Style struct {
	href string
}

func NewStyle() *Style {
	return new(Style)
}

func (self *Style) SetHref(href string) {
	self.href = href
}

func (self Style) String() string {
	var result string
	var attrs []string
	attrs = append(attrs, "rel=\"stylesheet\"")
	attrs = append(attrs, "type=\"text/css\"")
	attrs = append(attrs, fmt.Sprintf("href=\"%s\"", self.href))
	result = fmt.Sprintf("<link %s>", strings.Join(attrs, " "))
	return result
}
