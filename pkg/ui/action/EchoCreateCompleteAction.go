package action

import (
	"github.com/vit1251/golden/pkg/msg"
	"net/http"
)

type EchoCreateComplete struct {
	Action
}

func NewEchoCreateCompleteAction() *EchoCreateComplete {
	return new(EchoCreateComplete)
}

func (self *EchoCreateComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *msg.AreaManager
	self.Container.Invoke(func(am *msg.AreaManager) {
		areaManager = am
	})

	areaName := "RU.GOLDEN"

	areaManager.RemoveAreaByName(areaName)

}
