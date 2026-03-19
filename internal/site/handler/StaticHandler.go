package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

type StaticHandler struct {
}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse parameters */
	var filename string = r.PathValue("filename")

	/* Determine Content-Type */
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)

	/* Read content */
	resourcePath := path.Join("static", filename)
	log.Printf("Resource path %s", resourcePath)
	content, err1 := staticContent.ReadFile(resourcePath)
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
