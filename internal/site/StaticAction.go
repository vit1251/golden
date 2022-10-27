package site

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/internal/site/action"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

type StaticAction struct {
	action.Action
}

func NewStaticAction() *StaticAction {
	return new(StaticAction)
}

func (self *StaticAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse parameters */
	vars := mux.Vars(r)
	name := vars["name"]
	log.Printf("Request %s", name)

	/* Determine Content-Type */
	ext := filepath.Ext(name)
	contentType := mime.TypeByExtension(ext)

	/* Read content */
	resourcePath := path.Join("static", name)
	log.Printf("Resource path %s", resourcePath)
	content, err1 := staticContent.ReadFile(resourcePath)

	if err1 == nil {

		size := len(content)

		w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
		w.Header().Set("Content-Type", contentType)

		w.WriteHeader(200)

		image := bytes.NewReader(content)
		io.Copy(w, image)

	} else {

		log.Printf("Error: embed error %+v", err1)
		w.WriteHeader(404)

	}

}
