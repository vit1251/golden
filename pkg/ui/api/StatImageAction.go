package api

import (
	"encoding/json"
	stat2 "github.com/vit1251/golden/pkg/stat"
	"net/http"
)

type APIStatAction struct {
	APIAction
}

func NewAPIStatAction() *APIStatAction {
	sa := new(APIStatAction)
	return sa
}

func (self *APIStatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var statManager *stat2.StatManager
	self.Container.Invoke(func(sm *stat2.StatManager) {
		statManager = sm
	})

	/* Get statistics */
	sums, err1 := statManager.GetMessageSummary()
	if err1 != nil {
		panic(err1)
	}

	//
	js, err := json.Marshal(sums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
