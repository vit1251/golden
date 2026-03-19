package site2

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/vit1251/golden/internal/site2/api"
)

type PublicAction struct {
	api.Action
}

func NewPublicAction() *PublicAction {
	return &PublicAction{}
}

func (self *PublicAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse parameters */
	var name string = r.PathValue("name")
	log.Printf("Request %s", name)

	/* Determine Content-Type */
	ext := filepath.Ext(name)
	contentType := mime.TypeByExtension(ext)

	/* Read content */
	resourcePath := path.Join("public", name)
	log.Printf("Resource path %s", resourcePath)
	content, err1 := publicFS.ReadFile(resourcePath)

	if err1 != nil {
		log.Printf("Error: embed error %+v", err1)
		w.WriteHeader(404)
		return
	}

	size := len(content)

	w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(200)

	image := bytes.NewReader(content)
	io.Copy(w, image)

}
