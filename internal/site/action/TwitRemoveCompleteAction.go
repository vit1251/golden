package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type TwitRemoveCompleteAction struct {
	Action
}

func NewTwitRemoveCompleteAction() *TwitRemoveCompleteAction {
	return new(TwitRemoveCompleteAction)
}

func (self TwitRemoveCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	twitMapper := mapperManager.GetTwitMapper()

	/* Get "echoname" in user request */
	vars := mux.Vars(r)

	/* Restore Twit ID */
	twitId := vars["twitid"]

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
