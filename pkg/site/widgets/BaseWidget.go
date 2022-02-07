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

	/* Set scripts */
	mainScript := NewScript()
	mainScript.SetSrc("/static/custom.js")
	mainScript.SetDefer(true)
	bw.AddScript(*mainScript)

	/* Set theme style */
	mainStyle := NewStyle()
	mainStyle.SetHref("/static/custom.css")
	bw.AddStyle(*mainStyle)

	/* Set main style */
	themeStyle := NewStyle()
	themeStyle.SetHref("/static/theme_dark.css")
	bw.AddStyle(*themeStyle)

	/* Set main style */
	modernStyle := NewStyle()
	modernStyle.SetHref("/static/modern.css")
	bw.AddStyle(*modernStyle)

	/* Set print style */
	printStyle := NewStyle()
	printStyle.SetHref("/static/print.css")
	printStyle.SetMedia("print")
	bw.AddStyle(*printStyle)

	return bw
}

func (self *BaseWidget) AddStyle(s Style) *BaseWidget {
	self.styles = append(self.styles, s)
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

func (self *BaseWidget) AddScript(script Script) *BaseWidget {
	self.scripts = append(self.scripts, script)
	return self
}
