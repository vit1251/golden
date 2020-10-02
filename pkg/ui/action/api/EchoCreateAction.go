package api

import (
	"encoding/json"
	"net/http"
)

type EchoCreateAction struct {
	Action
}

func NewEchoCreateAction() *EchoCreateAction {
	return new(EchoCreateAction)
}

func (self *EchoCreateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	p := make(map[string]interface{})
	p["code"] = 1
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
