package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type FileEchoAreaUploadAction struct {
	Action
}

func NewFileEchoAreaUploadAction() *FileEchoAreaUploadAction {
	return new(FileEchoAreaUploadAction)
}

func (self FileEchoAreaUploadAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	configManager := self.restoreConfigManager()
	mapperManager := self.restoreMapperManager()
	fileAreaMapper := mapperManager.GetFileAreaMapper()

	/* Get BOSS address */
	newConfig := configManager.GetConfig()

	/* Get params */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
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

	container.AddWidget(containerVBox)

	vBox.Add(container)

	composeForm := widgets.NewFormWidget().
		SetAction(fmt.Sprintf("/file/%s/upload/complete", area.GetName())).
		SetEnctype("multipart/form-data").
		SetMethod("POST")

	/* Create form */
	newForm := widgets.NewVBoxWidget()

	/* File AP200830.ZIP */
	newForm.Add(widgets.NewFormFileInputWidget().SetTitle("File").SetName("file"))

	/* To */
	newForm.Add(widgets.NewFormInputWidget().SetTitle("To").SetName("to").SetValue(newConfig.Main.Link))

	/* Description */
	newForm.Add(widgets.NewFormInputWidget().SetTitle("Desc").SetName("desc"))

	/* Long description */
	newForm.Add(widgets.NewFormTextWidget().SetName("ldesc"))

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
