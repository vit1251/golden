package handler

import (
	"fmt"
	"net/http"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type TwitRemoveCompleteHandler struct {
	registry *registry.Container
}

func NewTwitRemoveCompleteHandler(registry *registry.Container) *TwitRemoveCompleteHandler {
	return &TwitRemoveCompleteHandler{
		registry: registry,
	}
}

func (self TwitRemoveCompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	twitMapper := mapperManager.GetTwitMapper()

	/* Restore Twit ID */
	var twitId string = r.PathValue("twitid")

	/* Remove by ID */
	err1 := twitMapper.RemoveById(twitId)
	if err1 != nil {
		status := fmt.Sprintf("Fail in RemoveById on twitMapper: err = %+v", err1)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/twit")
	http.Redirect(w, r, newLocation, 303)

}
