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

//

func (self *StaticAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	content := assets.Main[name]
	fmt.Printf("size = %+v\n", len(content) )

	ext := filepath.Ext(name)
	contnetType := mime.TypeByExtension(ext)
	fmt.Printf("contnetType = %+v\n", contnetType )
	w.Header().Set("Content-Type", contnetType)

	image := bytes.NewReader(content)

	size, err := io.Copy(w, image)
	fmt.Printf("static: size = %+v err = %+v\n", size, err)

}
