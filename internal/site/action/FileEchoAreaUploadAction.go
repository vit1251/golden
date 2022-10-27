package action

import (
	"fmt"
	"github.com/gorilla/mux"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/mapper"
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

	configManager := config.RestoreConfigManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
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
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	composeForm := widgets2.NewFormWidget().
		SetAction(fmt.Sprintf("/file/%s/upload/complete", area.GetName())).
		SetEnctype("multipart/form-data").
		SetMethod("POST")

	/* Create form */
	newForm := widgets2.NewVBoxWidget()

	/* File AP200830.ZIP */
	newForm.Add(widgets2.NewFormFileInputWidget().SetTitle("File").SetName("file"))

	/* To */
	newForm.Add(widgets2.NewFormInputWidget().SetTitle("To").SetName("to").SetValue(newConfig.Main.Link))

	/* Description */
	newForm.Add(widgets2.NewFormInputWidget().SetTitle("Desc").SetName("desc"))

	/* Long description */
	newForm.Add(widgets2.NewFormTextWidget().SetName("ldesc"))

	/* Complete */
	newForm.Add(widgets2.NewFormButtonWidget().SetType("submit").SetTitle("Send"))

	composeForm.SetWidget(newForm)

	containerVBox.Add(composeForm)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
