package site2

import (
	"bytes"
	"github.com/vit1251/golden/internal/site2/api"
	"net/http"
)

type IndexAction struct {
	api.Action
}

func NewIndexAction() *IndexAction {
	return new(IndexAction)
}

func (self *IndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    var cache bytes.Buffer
    cache.Grow(1024)

    cache.WriteString("<!DOCTYPE html>\n")
    cache.WriteString("<html>\n")
    cache.WriteString("<head>\n")
    cache.WriteString("    <link rel=\"stylesheet\" href=\"/public/main.css\">\n")
    cache.WriteString("    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
    cache.WriteString("    <title>Golden Point</title>\n")
    cache.WriteString("</head>\n")
    cache.WriteString("<body class=\"dark\">\n")
    cache.WriteString("    <div id=\"root\"></div>\n")
    cache.WriteString("    <script src=\"/public/main.js\"></script>\n")
    cache.WriteString("</body>\n")
    cache.WriteString("</html>\n")

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    cache.WriteTo(w)

}
