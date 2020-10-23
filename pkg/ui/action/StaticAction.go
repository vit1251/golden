package action

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/assets"
	"io"
	"mime"
	"net/http"
	"path/filepath"
)

type StaticAction struct {
	Action
}

func NewStaticAction() *StaticAction {
	return new(StaticAction)
}

func (self *StaticAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	if content, ok := assets.Main[name]; ok {
		ext := filepath.Ext(name)
		contentType := mime.TypeByExtension(ext)

		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
		w.Header().Set("Content-Type", contentType)

		w.WriteHeader(200)

		image := bytes.NewReader(content)
		io.Copy(w, image)

	} else {

		w.WriteHeader(404)

	}

}
