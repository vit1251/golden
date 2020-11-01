package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
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

	areaManager := self.restoreAreaManager()

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

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/remove", area.GetName())).
			SetIcon("icofont-remove").
			SetLabel("Remove echo")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/purge", area.GetName())).
			SetIcon("icofont-purge").
			SetLabel("Purge echo"))

	containerVBox.Add(amw)

	headerWidget := widgets.NewHeaderWidget().
		SetTitle(fmt.Sprintf("Settings on area %s", area.GetName()))

	containerVBox.Add(headerWidget)

	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/update/complete", area.GetName()))
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

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
