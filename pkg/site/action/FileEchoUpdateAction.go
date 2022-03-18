package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type FileEchoUpdateAction struct {
	Action
}

func NewFileEchoUpdateAction() *FileEchoUpdateAction {
	return new(FileEchoUpdateAction)
}

func (self FileEchoUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	//fileMapper := mapperManager.GetFileMapper()
	fileAreaMapper := mapperManager.GetFileAreaMapper()

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
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

	container.AddWidget(containerVBox)

	vBox.Add(container)

	/* Context actions */
	actionBar := self.renderActions(area)
	containerVBox.Add(actionBar)

	headerWidget := widgets.NewHeaderWidget().
		SetTitle(fmt.Sprintf("Settings on area %s", area.GetName()))

	containerVBox.Add(headerWidget)

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self FileEchoUpdateAction) renderActions(area *mapper.FileArea) widgets.IWidget {

	actionBar := widgets.NewActionMenuWidget()

	actionBar.Add(widgets.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/remove", area.GetName())).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Remove echo"))

	//actionBar.Add(widgets.NewMenuAction().
	//			SetLink(fmt.Sprintf("/file/%s/purge", area.GetName())).
	//			SetIcon("icofont-purge").
	//			SetLabel("Purge echo"))

	return actionBar

}
