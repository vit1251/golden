package action

import "net/http"

type EchoCreateAction struct {
	Action
}

func NewEchoCreateAction() *EchoCreateAction {
	return new(EchoCreateAction)
}

func (self *EchoCreateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
