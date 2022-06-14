package action

import (
	"github.com/vit1251/golden/pkg/mapper"
	"log"
	"net/http"
)

type EchoAreaCreateComplete struct {
	Action
}

func NewEchoAreaCreateCompleteAction() *EchoAreaCreateComplete {
	return new(EchoAreaCreateComplete)
}

func (self *EchoAreaCreateComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := self.restoreUrlManager()
	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	echoTag := r.Form.Get("echoname")
	log.Printf("echoTag = %v", echoTag)

	a := mapper.NewArea()
	a.SetName(echoTag)
	err2 := echoAreaMapper.Register(a)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	newAreaAddr := urlManager.CreateUrl("/echo/{area_index}").
		SetParam("area_index", a.GetAreaIndex()).
		Build()

	http.Redirect(w, r, newAreaAddr, 303)

}
