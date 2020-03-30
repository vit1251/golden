package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/tosser"
	"log"
	"net/http"
)

type EchoComposeCompleteAction struct {
	Action
}

func NewEchoComposeCompleteAction() (*EchoComposeCompleteAction) {
	cca := new(EchoComposeCompleteAction)
	return cca
}

func (self *EchoComposeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var tosserManager *tosser.TosserManager
	self.Container.Invoke(func(am *area2.AreaManager, tm *tosser.TosserManager) {
		areaManager = am
		tosserManager = tm
	})

	//
	vars := mux.Vars(r)
	//
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	//
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	to := r.Form.Get("to")
	subj := r.Form.Get("subject")
	body := r.Form.Get("body")
	log.Printf("to = %s subj = %s body = %s", to, subj, body)

	/* Create message */
	em := tosserManager.NewEchoMessage()
	em.Subject = subj
	em.Body = body
	em.AreaName = area.Name()
	em.To = to

	/* Delivery message */
	err3 := tosserManager.WriteEchoMessage(em)
	if err3 != nil {
		panic(err3)
	}

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
