package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)


type FileAreaUploadAction struct {
	Action
}

func NewFileAreaUploadAction() *FileAreaUploadAction {
	return new(FileAreaUploadAction)
}

func (self FileAreaUploadAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fileManager := self.restoreFileManager()

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := fileManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

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
		SetAction(fmt.Sprintf("/file/%s/compose/complete", area.GetName())).
		SetMethod("POST")

	/* Create form */
	newForm := widgets.NewVBoxWidget()

	/* File AP200830.ZIP */
	newForm.Add(widgets.NewFormFileInputWidget().SetTitle("File").SetName("file"))

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
