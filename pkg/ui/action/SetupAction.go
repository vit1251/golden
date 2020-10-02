package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type SetupAction struct {
	Action
}

func NewSetupAction() (*SetupAction) {
	sa := new(SetupAction)
	return sa
}

func (self *SetupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Setup manager operation */
	var setupManager *setup.ConfigManager
	self.Container.Invoke(func(sm *setup.ConfigManager) {
		setupManager = sm
	})

	params := setupManager.GetParams()
	log.Printf("params = %+v", params)

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	setupForm := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction("/setup/complete")

	/* Add custom param field */
	setupFormBox := widgets.NewVBoxWidget()
	for _, param := range params {
		log.Printf("param = %+v", param)
		self.createInputField(setupFormBox,
			param.Name,
			param.Summary,
			param.Value)
	}
	setupFormBox.Add(widgets.NewFormButtonWidget().SetTitle("Save").SetType("submit"))
	setupForm.SetWidget(setupFormBox)

	containerVBox.Add(setupForm)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *SetupAction) createInputField(box *widgets.VBoxWidget, name string, summary string, value string) {

	mainDiv := widgets.NewDivWidget().
		SetClass("form-group row")

	mainDivBox := widgets.NewVBoxWidget()
	mainDiv.SetWidget(mainDivBox)

	mainTitle := widgets.NewDivWidget().
		SetClass("col-sm-2 col-form-label").
		SetContent(name)

	mainDivBox.Add(mainTitle)

	mainInput := widgets.NewFormInputWidget().
		SetTitle(summary).
		SetName(name).
		SetValue(value)

	mainDivBox.Add(mainInput)

	box.Add(mainDiv)

}
