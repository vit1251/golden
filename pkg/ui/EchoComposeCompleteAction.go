package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/common"
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

	master := common.GetMaster()

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
	areaManager := master.AreaManager
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %v", area)

	//
	to := r.Form.Get("to")
	subj := r.Form.Get("subject")
	body := r.Form.Get("body")
	log.Printf("to = %s subj = %s body = %s", to, subj, body)

	/* Create message */
	em := master.TosserManager.NewEchoMessage()
	em.Subject = subj
	em.Body = body
	em.AreaName = area.Name
	em.To = to

	/* Delivery message */
	err3 := master.TosserManager.WriteEchoMessage(em)
	if err3 != nil {
		panic(err3)
	}

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
