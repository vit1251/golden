package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/tosser"
	"log"
	"net/http"
)

type EchoComposeCompleteAction struct {
	Action
}

func NewEchoComposeCompleteAction() *EchoComposeCompleteAction {
	cca := new(EchoComposeCompleteAction)
	return cca
}

func (self *EchoComposeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	tosserManager := self.restoreTosserManager()
	statMapper := mapperManager.GetStatMapper()

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
	area, err1 := echoAreaMapper.GetAreaByName(echoTag)
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
	em := tosser.NewEchoMessage()
	em.Subject = subj
	em.SetBody(body)
	em.AreaName = area.GetName()
	em.To = to

	/* Delivery message */
	err3 := tosserManager.WriteEchoMessage(em)
	if err3 != nil {
		panic(err3)
	}

	/* Register packet */
	if err := statMapper.RegisterOutPacket(); err != nil {
		log.Printf("Fail on RegisterInPacket: err = %+v", err)
	}
	if err := statMapper.RegisterOutMessage(); err != nil {
		log.Printf("Fail on RegisterOutMessage: err = %+v", err)
	}

	//
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)
}
