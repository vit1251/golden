package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type EchoCreateComplete struct {
	Action
}

func NewEchoCreateCompleteAction() *EchoCreateComplete {
	return new(EchoCreateComplete)
}

func (self *EchoCreateComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	echoTag := r.Form.Get("echoname")
	fmt.Printf("echoTag = %v", echoTag)

	a := mapper.NewArea()
	a.SetName(echoTag)
	echoAreaMapper.Register(a)

	//
	newLocation := fmt.Sprintf("/echo/%s", a.GetName())
	http.Redirect(w, r, newLocation, 303)

}
