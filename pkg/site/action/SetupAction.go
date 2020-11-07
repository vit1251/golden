package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
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

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()

	params := configMapper.GetParams()
	log.Printf("params = %+v", params)

	/* Render */
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

	setupForm := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction("/setup/complete")

	/* Add custom param field */
	setupFormBox := widgets.NewVBoxWidget()
	for _, param := range params {
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
		SetClass("").
		SetContent(name)

	mainDivBox.Add(mainTitle)

	mainInput := widgets.NewFormInputWidget().
		SetTitle(summary).
		SetName(name).
		SetValue(value)

	mainDivBox.Add(mainInput)

	box.Add(mainDiv)

}
