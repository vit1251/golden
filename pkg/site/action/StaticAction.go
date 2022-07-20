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

        /* Parse parameters */
	vars := mux.Vars(r)
	name := vars["name"]

        /* Determine Content-Type */
	ext := filepath.Ext(name)
	contentType := mime.TypeByExtension(ext)

        /* Read content */
	content, err1 := assets.Content.ReadFile(name)
        if err1 == nil {

                size := len(content)

		w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
		w.Header().Set("Content-Type", contentType)

		w.WriteHeader(200)

		image := bytes.NewReader(content)
		io.Copy(w, image)

	} else {

		w.WriteHeader(404)

	}

}
