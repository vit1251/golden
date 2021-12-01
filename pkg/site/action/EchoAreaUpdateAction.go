package action

import (
	"fmt"
	"github.com/gorilla/mux"
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

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := echoAreaMapper.GetAreaByName(echoTag)
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

	container.SetWidget(containerVBox)

	vBox.Add(container)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/remove", area.GetName())).
			SetIcon("icofont-remove").
			SetLabel("Remove echo")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/purge", area.GetName())).
			SetIcon("icofont-purge").
			SetLabel("Purge echo"))

	containerVBox.Add(amw)

	headerWidget := widgets.NewHeaderWidget().
		SetTitle(fmt.Sprintf("Settings on area %s", area.GetName()))

	containerVBox.Add(headerWidget)

	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/update/complete", area.GetName()))
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
