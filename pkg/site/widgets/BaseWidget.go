package widgets

import (
	"fmt"
	"io"
)

type BaseWidget struct {
	mainWidget IWidget
	styles     []Style
	scripts    []Script
}

func NewBaseWidget() *BaseWidget {
	bw := new(BaseWidget)
	bw.AddStyle("/static/custom.css")
	bw.AddScript("/static/custom.js", true)
	bw.AddStyle("/assets/css/main.css")
	return bw
}

func (self *BaseWidget) AddStyle(path string) *BaseWidget {
	s := NewStyle()
	s.SetHref(path)
	self.styles = append(self.styles, *s)
	return self
}

func (self *BaseWidget) Render(w io.Writer) error {

	builder := NewByteBuilder()

	builder.AppendString("<!DOCTYPE html>\n")
	builder.AppendString("<html>\n")
	builder.AppendString("<head>\n")
	builder.AppendString("\t<meta charset=\"utf-8\">\n")
	builder.AppendString("\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n")
	builder.AppendString("\t<title>Golden Point</title>\n")
	for _, s := range self.styles {
		builder.AppendString(fmt.Sprintf("\t%s\n", s.String()))
	}
	for _, script := range self.scripts {
		builder.AppendString(fmt.Sprintf("\t%s\n", script.String()))
	}
	builder.AppendString("</head>\n")
	builder.AppendString("<body class=\"dark\">\n")

	self.mainWidget.Render(builder)

	builder.AppendString("</body>\n")
	builder.AppendString("</html>\n")

	content := builder.Byte()
	_, err1 := w.Write(content)

	return err1
}

func (self *BaseWidget) SetWidget(widget IWidget) {
	self.mainWidget = widget
}

func (self *BaseWidget) AddScript(src string, defered bool) *BaseWidget {

	script := NewScript()
	script.Src = src
	script.Defer = defered
	self.scripts = append(self.scripts, *script)

	return self
}
