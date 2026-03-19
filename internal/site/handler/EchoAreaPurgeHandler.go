package handler

import (
	"fmt"
	"log"
	"net/http"

	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type EchoAreaPurgeHandler struct {
	registry *registry.Container
}

func NewEchoAreaPurgeHandler(registry *registry.Container) *EchoAreaPurgeHandler {
	return &EchoAreaPurgeHandler{
		registry: registry,
	}
}

func (self *EchoAreaPurgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.registry)
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Parse URL parameters */
	var areaIndex string = r.PathValue("echoname")
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

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
		SetTitle("Purge area?")
	containerVBox.Add(headerWidget)

	//fmt.Sprintf(, area.GetName())
	purgeCompleteAddr := urlManager.CreateUrl("/echo/{area_index}/purge/complete").
		SetParam("area_index", area.GetAreaIndex()).
		Build()

	formWidget := widgets2.NewFormWidget().
		SetMethod("POST").
		SetAction(purgeCompleteAddr)
	formVBox := widgets2.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	qustionWidget := widgets2.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to remove '%s' area?", area.GetName()))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets2.NewFormButtonWidget().
		SetTitle("Purge")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
