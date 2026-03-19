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

type EchoAreaUpdateHandler struct {
	registry *registry.Container
}

func NewEchoAreaUpdateHandler(registry *registry.Container) *EchoAreaUpdateHandler {
	return &EchoAreaUpdateHandler{
		registry: registry,
	}
}

func (self *EchoAreaUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.registry)
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	//
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

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	/* Context handlers */
	handlersBar := self.renderHandlers(area)
	containerVBox.Add(handlersBar)

	headerWidget := widgets2.NewHeaderWidget().
		SetTitle(fmt.Sprintf("Settings on area %s", area.GetName()))

	containerVBox.Add(headerWidget)

	settingsSaveAddr := urlManager.CreateUrl("/echo/{area_index}/update/complete").
		SetParam("area_index", area.GetAreaIndex()).
		Build()

	formWidget := widgets2.NewFormWidget().
		SetMethod("POST").
		SetAction(settingsSaveAddr)
	formVBox := widgets2.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	/* Summary */
	summaryInputWidget := widgets2.NewFormInputWidget().
		SetTitle("Summary").SetName("summary").SetValue(area.Summary)

	formVBox.Add(summaryInputWidget)

	/* Charset */
	chrsInputWidget := widgets2.NewFormInputWidget().
		SetTitle("Charset").SetName("charset").SetValue(area.GetCharset())
	formVBox.Add(chrsInputWidget)

	/* Charset */
	orderInputWidget := widgets2.NewFormInputWidget().
		SetTitle("Sort order").SetName("order").SetValue(fmt.Sprintf("%d", area.GetOrder()))
	formVBox.Add(orderInputWidget)

	/* Save button */
	btnWidget := widgets2.NewFormButtonWidget().
		SetTitle("Save")
	formVBox.Add(btnWidget)

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *EchoAreaUpdateHandler) renderHandlers(area *mapper.Area) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.registry)
	actionBar := widgets2.NewActionMenuWidget()

	/* Remove area handler button */
	removeAreaAddr := urlManager.CreateUrl("/echo/{area_index}/remove").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(removeAreaAddr).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Remove echo"))

	/* Purge area handler button */
	purgeAreaAddr := urlManager.CreateUrl("/echo/{area_index}/purge").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(purgeAreaAddr).
		SetIcon("icofont-purge").
		SetClass("mr-2").
		SetLabel("Purge echo"))

	return actionBar
}
