package action

import (
	"encoding/json"
	"net/http"
)

type EchoCreateApiAction struct {
	Action
}

func NewEchoCreateApiAction() *EchoCreateApiAction {
	return new(EchoCreateApiAction)
}

func (self *EchoCreateApiAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	p := make(map[string]interface{})
	p["code"] = 1
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
