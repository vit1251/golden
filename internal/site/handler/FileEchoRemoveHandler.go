package handler

import (
	"fmt"
	"log"
	"net/http"

	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type FileEchoRemoveHandler struct {
	registry *registry.Container
}

func NewFileEchoRemoveHandler(registry *registry.Container) *FileEchoRemoveHandler {
	return &FileEchoRemoveHandler{
		registry: registry,
	}
}

func (self FileEchoRemoveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	//fileMapper := mapperManager.GetFileMapper()
	fileAreaMapper := mapperManager.GetFileAreaMapper()

	//
	var echoTag string = r.PathValue("echoname")
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Render question */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets2.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")
	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()
	container.AddWidget(containerVBox)

	//<h1>Delete message?</h1>
	headerWidget := widgets2.NewHeaderWidget().
		SetTitle("Delete area?")
	containerVBox.Add(headerWidget)

	//
	formWidget := widgets2.NewFormWidget().
		SetMethod("POST").
		SetAction(fmt.Sprintf("/file/%s/remove/complete", area.GetName()))
	formVBox := widgets2.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	qustionWidget := widgets2.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to remove '%s' area?", area.GetName()))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets2.NewFormButtonWidget().
		SetTitle("Remove")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
