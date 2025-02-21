package site2

import (
	"bytes"
	"fmt"
	"github.com/vit1251/golden/internal/site2/api"
	"io"
	"net/http"
)

type IndexAction struct {
	api.Action
}

func NewIndexAction() *IndexAction {
	return new(IndexAction)
}

func (self *IndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    out := &bytes.Buffer{}

    fmt.Fprintf(out, "<!DOCTYPE html>")
    fmt.Fprintf(out, "<html lang=\"en\">")
    fmt.Fprintf(out, "<head>")
    fmt.Fprintf(out, "<link rel=\"stylesheet\" href=\"/public/main.css\">")
    fmt.Fprintf(out, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
    fmt.Fprintf(out, "<title>Golden Point</title>")
    fmt.Fprintf(out, "</head>")
    fmt.Fprintf(out, "<body class=\"dark\">")
    fmt.Fprintf(out, "<div id=\"root\"></div>")
    fmt.Fprintf(out, "<script src=\"/public/main.js\"></script>")
    fmt.Fprintf(out, "</body>")
    fmt.Fprintf(out, "</html>")

    content := out.Bytes()

    size := len(content)

    w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
    w.Header().Set("Content-Type", "text/html")

    w.WriteHeader(200)

    image := bytes.NewReader(content)
    io.Copy(w, image)

}
