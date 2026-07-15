package handler

import (
    "log"
    "net/http"

    "github.com/vit1251/golden/internal/um"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoAreaCreateComplete struct {
    registry *registry.Container
}

func NewEchoAreaCreateCompleteHandler(registry *registry.Container) *EchoAreaCreateComplete {
    return &EchoAreaCreateComplete{
	registry: registry,
    }
}

func (self *EchoAreaCreateComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    urlManager := um.RestoreUrlManager(self.registry)
    mapperManager := mapper.RestoreMapperManager(self.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()

    err := r.ParseForm()
    if err != nil {
	panic(err)
    }

    //
    echoTag := r.Form.Get("echoname")
    log.Printf("echoTag = %v", echoTag)

    a := mapper.NewArea()
    a.SetName(echoTag)
    err2 := echoAreaMapper.Register(a)
    if err2 != nil {
	panic(err2)
    }

    /* Redirect */
    newAreaAddr := urlManager.CreateUrl("/echo/{area_index}").
	SetParam("area_index", a.GetAreaIndex()).
	Build()

    http.Redirect(w, r, newAreaAddr, 303)

}
