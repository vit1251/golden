package site2

import (
	"bytes"
	"fmt"
	"github.com/vit1251/golden/pkg/site2/api"
	"io"
	"log"
	"path"
	"net/http"
)

type IndexAction struct {
	api.Action
}

func NewIndexAction() *IndexAction {
	return new(IndexAction)
}

func (self *IndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Read content */
	resourcePath := path.Join("public", "index.html")
	content, err1 := publicContent.ReadFile(resourcePath)

        if err1 == nil {

                size := len(content)

		w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
		w.Header().Set("Content-Type", "text/html")

		w.WriteHeader(200)

		image := bytes.NewReader(content)
		io.Copy(w, image)

	} else {

                log.Printf("Error: embed error %+v", err1)
		w.WriteHeader(404)

	}

}
