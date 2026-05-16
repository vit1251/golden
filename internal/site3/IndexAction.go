package site3

import (
	"bytes"
	"net/http"

	"github.com/vit1251/golden/pkg/registry"
)

type IndexHandler struct {
	registry *registry.Container
}

func NewIndexAction(registry *registry.Container) *IndexHandler {
	return &IndexHandler{
		registry: registry,
	}
}

func (self *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var cache bytes.Buffer
	cache.Grow(1024)

	cache.WriteString("<!DOCTYPE html>\n")
	cache.WriteString("<html>\n")
	cache.WriteString("<head>\n")
	cache.WriteString("    <link rel=\"stylesheet\" href=\"/public/bundle.css\">\n")
	cache.WriteString("    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	cache.WriteString("    <title>Golden Point</title>\n")
	cache.WriteString("</head>\n")
	cache.WriteString("<body>\n")
	cache.WriteString("    <div id=\"app\"></div>\n")
	cache.WriteString("    <script src=\"/public/bundle.js\"></script>\n")
	cache.WriteString("</body>\n")
	cache.WriteString("</html>\n")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	cache.WriteTo(w)

}
