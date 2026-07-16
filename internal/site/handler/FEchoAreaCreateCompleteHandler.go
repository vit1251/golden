package handler

import (
    "net/http"

    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type FEchoAreaCreateCompleteHandler struct {
    registry *registry.Container
}

func NewFEchoAreaCreateCompleteHandler(registry *registry.Container) *FEchoAreaCreateCompleteHandler {
    return &FEchoAreaCreateCompleteHandler{
    	registry: registry,
    }
}

func (h *FEchoAreaCreateCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    fileAreaMapper := mapperManager.GetFileAreaMapper()

    err1 := r.ParseForm()
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
	return
    }
    echoTag := r.Form.Get("fileecho")

    /* Create File area */
    a := mapper.NewFileArea()
    a.SetName(echoTag)
    err2 := fileAreaMapper.CreateFileArea(a)
    if err2 != nil {
        http.Error(w, err2.Error(), http.StatusInternalServerError)
	return
    }

    http.Redirect(w, r, "/file/" + a.GetName(), 303)

}
