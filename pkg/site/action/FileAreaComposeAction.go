package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type FileAreaComposeAction struct {
	Action
}

func NewFileAreaComposeAction() *FileAreaComposeAction {
	return new(FileAreaComposeAction)
}

func (self *FileAreaComposeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fileArea := "NASA"

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

	composeForm := widgets.NewFormWidget().
		SetAction(fmt.Sprintf("/file/%s/compose/complete", fileArea)).
		SetMethod("POST")

	/* Create form */
	newForm := widgets.NewVBoxWidget()

	/* File AP200830.ZIP */
	newForm.Add(widgets.NewFormInputWidget().SetTitle("File").SetName("file"))

	/* Desc NASA Astronomy Picture of the Day (plus published report) */
	newForm.Add(widgets.NewFormInputWidget().SetTitle("Desc").SetName("desc"))

	/* Complete */
	newForm.Add(widgets.NewFormButtonWidget().SetType("submit").SetTitle("Send"))

	composeForm.SetWidget(newForm)

	containerVBox.Add(composeForm)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
