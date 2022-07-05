package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
	"github.com/vit1251/golden/pkg/um"
	"log"
	"net/http"
)

type EchoAreaPurgeAction struct {
	Action
}

func NewEchoAreaPurgeAction() *EchoAreaPurgeAction {
	ra := new(EchoAreaPurgeAction)
	return ra
}

func (self *EchoAreaPurgeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Parse URL parameters */
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
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.AddWidget(containerVBox)

	//<h1>Delete message?</h1>
	headerWidget := widgets.NewHeaderWidget().
		SetTitle("Purge area?")
	containerVBox.Add(headerWidget)

	//fmt.Sprintf(, area.GetName())
	purgeCompleteAddr := urlManager.CreateUrl("/echo/{area_index}/purge/complete").
		SetParam("area_index", area.GetAreaIndex()).
		Build()

	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(purgeCompleteAddr)
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	qustionWidget := widgets.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to remove '%s' area?", area.GetName()))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets.NewFormButtonWidget().
		SetTitle("Purge")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
