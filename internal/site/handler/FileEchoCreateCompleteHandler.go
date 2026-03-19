package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type FileEchoCreateComplete struct {
	registry *registry.Container
}

func NewFileEchoCreateCompleteHandler(registry *registry.Container) *FileEchoCreateComplete {
	return &FileEchoCreateComplete{
		registry: registry,
	}
}

func (self *FileEchoCreateComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Prepare mappers */
	mapperManager := mapper.RestoreMapperManager(self.registry)
	fileAreaMapper := mapperManager.GetFileAreaMapper()

	/* Parse POST parameters */
	err1 := r.ParseForm()
	if err1 != nil {
		panic(err1)
	}
	echoTag := r.Form.Get("fileecho")
	log.Printf("echoTag = %v", echoTag)

	/* Create File area */
	a := mapper.NewFileArea()
	a.SetName(echoTag)
	err2 := fileAreaMapper.CreateFileArea(a)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect user */
	newLocation := fmt.Sprintf("/file/%s", a.GetName())
	http.Redirect(w, r, newLocation, 303)

}
