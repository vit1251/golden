package action

import (
	"fmt"
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

	areaManager := self.restoreAreaManager()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	echoTag := r.Form.Get("echoname")
	fmt.Printf("echoTag = %v", echoTag)

	a := msg.NewArea()
	a.SetName(echoTag)
	areaManager.Register(a)

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)

}
