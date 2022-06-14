package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type EchoAreaUpdateAction struct {
	Action
}

func NewEchoAreaUpdateAction() *EchoAreaUpdateAction {
	ea := new(EchoAreaUpdateAction)
	return ea
}

func (self *EchoAreaUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := self.restoreUrlManager()
	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	//
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
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
	actionsBar := self.renderActions(area)
	containerVBox.Add(actionsBar)

	headerWidget := widgets.NewHeaderWidget().
		SetTitle(fmt.Sprintf("Settings on area %s", area.GetName()))

	containerVBox.Add(headerWidget)

	settingsSaveAddr := urlManager.CreateUrl("/echo/{area_index}/update/complete").
		SetParam("area_index", area.GetAreaIndex()).
		Build()

	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(settingsSaveAddr)
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	/* Summary */
	summaryInputWidget := widgets.NewFormInputWidget().
		SetTitle("Summary").SetName("summary").SetValue(area.Summary)

	formVBox.Add(summaryInputWidget)

	/* Charset */
	chrsInputWidget := widgets.NewFormInputWidget().
		SetTitle("Charset").SetName("charset").SetValue(area.GetCharset())
	formVBox.Add(chrsInputWidget)

	/* Charset */
	orderInputWidget := widgets.NewFormInputWidget().
		SetTitle("Sort order").SetName("order").SetValue(fmt.Sprintf("%d", area.GetOrder()))
	formVBox.Add(orderInputWidget)

	/* Save button */
	btnWidget := widgets.NewFormButtonWidget().
		SetTitle("Save")
	formVBox.Add(btnWidget)

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *EchoAreaUpdateAction) renderActions(area *mapper.Area) widgets.IWidget {

	urlManager := self.restoreUrlManager()
	actionBar := widgets.NewActionMenuWidget()

	/* Remove area action button */
	removeAreaAddr := urlManager.CreateUrl("/echo/{area_index}/remove").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	actionBar.Add(widgets.NewMenuAction().
		SetLink(removeAreaAddr).
		SetIcon("icofont-remove").
		SetClass("mr-2").
		SetLabel("Remove echo"))

	/* Purge area action button */
	purgeAreaAddr := urlManager.CreateUrl("/echo/{area_index}/purge").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	actionBar.Add(widgets.NewMenuAction().
		SetLink(purgeAreaAddr).
		SetIcon("icofont-purge").
		SetClass("mr-2").
		SetLabel("Purge echo"))

	return actionBar
}
