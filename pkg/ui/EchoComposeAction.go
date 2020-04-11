package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoComposeAction struct {
	Action
}

func NewEchoComposeAction() *EchoComposeAction {
	ca := new(EchoComposeAction)
	return ca
}

func (self *EchoComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	self.Container.Invoke(func(am *area2.AreaManager) {
		areaManager = am
	})

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Search echo area */

	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	formVBox := widgets.NewVBoxWidget()

	formWidget := widgets.NewFormWidget()
	formWidget.
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/message/compose/complete", area.Name())).
		SetWidget(formVBox)

	formVBox.Add(widgets.NewFormInputWidget().SetTitle("TO").SetName("to"))
	formVBox.Add(widgets.NewFormInputWidget().SetTitle("SUBJ").SetName("subject"))
	formVBox.Add(widgets.NewFormTextWidget().SetName("body"))
	formVBox.Add(widgets.NewFormButtonWidget().SetTitle("Compose").SetType("submit"))

	vBox.Add(formWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
