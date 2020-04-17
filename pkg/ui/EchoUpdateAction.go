package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoUpdateAction struct {
	Action
}

func NewEchoUpdateAction() *EchoUpdateAction {
	ea := new(EchoUpdateAction)
	return ea
}

func (self *EchoUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	self.Container.Invoke(func(am *area2.AreaManager) {
		areaManager = am
	})

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/remove", area.Name())).
			SetIcon("icofont-remove").
			SetLabel("Remove echo"))

	vBox.Add(amw)

	headerWidget := widgets.NewHeaderWidget().
		SetTitle(fmt.Sprintf("Settings on area %s", area.Name()))

	vBox.Add(headerWidget)

	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/update/complete", area.Name()))
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	vBox.Add(formWidget)

	inputWidget := widgets.NewFormInputWidget().
		SetTitle("Summary").SetName("summary").SetValue(area.Summary)

	formVBox.Add(inputWidget)

	btnWidget := widgets.NewFormButtonWidget().
		SetTitle("Save")
	formVBox.Add(btnWidget)

	//
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
