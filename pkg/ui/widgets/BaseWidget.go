package widgets

import "net/http"

type BaseWidget struct {
	mainWidget IWidget
}

func NewBaseWidget() *BaseWidget {
	bw := new(BaseWidget)
	return bw
}

func (self *BaseWidget) Render(w http.ResponseWriter) error {

	w.Write([]byte("<!DOCTYPE html>\n"))
	w.Write([]byte("<html>\n"))
	w.Write([]byte("<head>\n"))
	w.Write([]byte("\t<meta charset=\"utf-8\">\n"))
	w.Write([]byte("\t<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n"))
	w.Write([]byte("\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n"))
	w.Write([]byte("\t<title>Golden Point</title>\n"))
	w.Write([]byte("\t<link rel=\"stylesheet\" type=\"text/css\" href=\"/static/css/icofont.css\">\n"))
	w.Write([]byte("\t<link rel=\"stylesheet\" type=\"text/css\" href=\"/static/css/main.css\">\n"))
	w.Write([]byte("\t<link rel=\"stylesheet\" type=\"text/css\" href=\"/static/css/custom.css\">\n"))
	w.Write([]byte("\t<script src=\"/static/js/mousetrap-1.6.5/mousetrap.js\"></script>\n"))
	w.Write([]byte("\t<script src=\"/static/js/jquery-3.4.1/jquery-3.4.1.min.js\"></script>\n"))
	w.Write([]byte("\t<script src=\"/static/js/custom.js\"></script>\n"))
	w.Write([]byte("</head>\n"))
	w.Write([]byte("<body class=\"dark\">\n"))

	self.mainWidget.Render(w)

	w.Write([]byte("</body>\n"))
	w.Write([]byte("</html>\n"))

	return nil
}

func (self *BaseWidget) SetWidget(widget IWidget) {
	self.mainWidget = widget
}
