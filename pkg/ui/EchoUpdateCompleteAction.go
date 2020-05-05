package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
	"log"
	"net/http"
)

type EchoUpdateCompleteAction struct {
	Action
}

func NewEchoUpdateCompleteAction() *EchoUpdateCompleteAction {
	euc := new(EchoUpdateCompleteAction)
	return euc
}

func (self *EchoUpdateCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Parse POST parameters */
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	/* ... */
	var areaManager *msg.AreaManager
	self.Container.Invoke(func(am *msg.AreaManager) {
		areaManager = am
	})

	/* ... */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* ... */
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/**/
	area.Summary = r.PostForm.Get("summary")

	/* ... */
	err2 := areaManager.Update(area)
	if err2 != nil {
		panic(err2)
	}

	/* Render */
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
