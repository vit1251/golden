package widgets

import (
	"fmt"
	"net/http"
)

type style struct {
	Path string
}

type JSScipt struct {
	Path string
}

type BaseWidget struct {
	mainWidget IWidget
	styles     []*style
	scripts    []*JSScipt
}

func NewBaseWidget() *BaseWidget {
	bw := new(BaseWidget)
	bw.AddStyle("/static/css/main.css")
	bw.AddStyle("/static/css/custom.css")
	bw.AddScript("/static/js/custom.js")
	return bw
}

func (self *BaseWidget) AddStyle(path string) *BaseWidget {
	s := new(style)
	s.Path = path
	self.styles = append(self.styles, s)
	return self
}

func (self *BaseWidget) Render(w http.ResponseWriter) error {

	w.Write([]byte("<!DOCTYPE html>\n"))
	w.Write([]byte("<html>\n"))

	/* Headers */
	w.Write([]byte("<head>\n"))
	w.Write([]byte("\t<meta charset=\"utf-8\">\n"))
	w.Write([]byte("\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n"))

	w.Write([]byte("\t<title>Golden Point</title>\n"))

	for _, s := range self.styles {
		msg := fmt.Sprintf("\t<link rel=\"stylesheet\" type=\"text/css\" href=\"%s\">\n", s.Path)
		w.Write([]byte(msg))
	}

	for _, script := range self.scripts {
		msg := fmt.Sprintf("\t<script src=\"%s\"></script>\n", script.Path)
		w.Write([]byte(msg))
	}

	w.Write([]byte("</head>\n"))

	/* Body */
	w.Write([]byte("<body class=\"dark\">\n"))
	self.mainWidget.Render(w)
	w.Write([]byte("</body>\n"))

	w.Write([]byte("</html>\n"))

	return nil
}

func (self *BaseWidget) SetWidget(widget IWidget) {
	self.mainWidget = widget
}

func (self *BaseWidget) AddScript(s string) *BaseWidget {
	script := new(JSScipt)
	script.Path = s
	self.scripts = append(self.scripts, script)
	return self
}
