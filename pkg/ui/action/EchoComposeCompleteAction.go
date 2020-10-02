package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/stat"
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

	var areaManager *msg.AreaManager
	var tosserManager *tosser.TosserManager
	var statManager *stat.StatManager
	self.Container.Invoke(func(am *msg.AreaManager, tm *tosser.TosserManager, sm *stat.StatManager) {
		areaManager = am
		tosserManager = tm
		statManager = sm
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

	/* Register packet */
	if err := statManager.RegisterOutPacket(); err != nil {
		log.Printf("Fail on RegisterInPacket: err = %+v", err)
	}
	if err := statManager.RegisterOutMessage(); err != nil {
		log.Printf("Fail on RegisterOutMessage: err = %+v", err)
	}

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
