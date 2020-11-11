package widgets

import (
	"fmt"
	"strings"
)

type Script struct {
	Src string
	Async bool
	Defer bool
//	Type string
}

func NewScript() *Script {
	return new(Script)
}

func (self Script) String() string {
	var result string
	var attrs []string
	if self.Src != "" {
		src := fmt.Sprintf("src=\"%s\"", self.Src)
		attrs = append(attrs, src)
	}
	if self.Async {
		attrs = append(attrs, "async")
	}
	if self.Defer {
		attrs = append(attrs, "defer")
	}
	result = fmt.Sprintf("<script %s></script>", strings.Join(attrs, " "))
	return result
}

func (self Script) SetSrc(s string) {
	self.Src = s
}

func (self Script) SetDefer(b bool) {
	self.Defer = b
}
