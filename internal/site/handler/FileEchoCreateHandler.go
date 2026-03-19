package handler

import (
	"fmt"
	"net/http"

	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/registry"
)

type FileEchoCreateHandler struct {
	registry *registry.Container
}

func NewFileEchoCreateHandler(registry *registry.Container) *FileEchoCreateHandler {
	return &FileEchoCreateHandler{
		registry: registry,
	}
}

func (self *FileEchoCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Render */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets2.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	setupForm := widgets2.NewFormWidget().
		SetMethod("POST").
		SetAction("/file/create/complete")

	/* Add custom param field */
	setupFormBox := widgets2.NewVBoxWidget()

	self.createInputField(setupFormBox, "fileecho", "File area name", "?")

	setupFormBox.Add(widgets2.NewFormButtonWidget().SetTitle("Save").SetType("submit"))
	setupForm.SetWidget(setupFormBox)

	containerVBox.Add(setupForm)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *FileEchoCreateHandler) createInputField(box *widgets2.VBoxWidget, name string, summary string, value string) {

	mainDiv := widgets2.NewDivWidget().
		SetClass("form-group row")

	mainDivBox := widgets2.NewVBoxWidget()
	mainDiv.AddWidget(mainDivBox)

	mainTitle := widgets2.NewDivWidget().
		SetClass("col-sm-2 col-form-label").
		SetContent(name)

	mainDivBox.Add(mainTitle)

	mainInput := widgets2.NewFormInputWidget().
		SetTitle(summary).
		SetName(name).
		SetValue(value)

	mainDivBox.Add(mainInput)

	box.Add(mainDiv)

}
